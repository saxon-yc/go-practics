package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v2"
	"gorm.io/gorm/logger"
)

type ChartArgs struct {
	ReleaseName  string
	ChartName    string
	ChartVersion string
	NameSpace    string
	Values       string
	RepoName     string
	RepoURL      string
}
type ChartValues map[string]map[string]interface{}

func main() {
	UpdateComponet()
}

func UpdateComponet() (err error) {
	var isExisted bool
	// 包装参数
	chart := new(ChartArgs)
	if err = generateParameters(chart); err != nil {
		return err
	}

	if isExisted, err = releaseExists(chart.ReleaseName, chart.NameSpace); err != nil {
		return err
	}

	if isExisted { // 合并生成新的chart values
		if err = mergeChartValues(chart); err != nil {
			return err
		}

		// 卸载当前 release
		if err = uninstallRelease(chart); err != nil {
			return err
		}
	}

	// 添加和更新Helm 仓库
	if err = addHelmRepo(chart); err != nil {
		return err
	}

	// 安装新的 Helm chart
	if err = installNewChart(chart, isExisted); err != nil {
		return err
	}

	// TODO: 安装成功后更新 config-simple.yaml 文件

	log.Default().Print("Install successfully")

	return nil
}

func generateParameters(chart *ChartArgs) error {

	chart.ReleaseName = "ccm-qingcloud"
	chart.ChartName = "stable/ccm-qingcloud"
	chart.ChartVersion = "v0.1.0"
	chart.NameSpace = "kube-system"
	chart.Values = ""
	chart.RepoName = "stable" // 从数据库获取
	chart.RepoURL = "https://charts.kubesphere.io/stable"

	return nil
}

// uninstallRelease 卸载 Helm release
func uninstallRelease(chart *ChartArgs) error {
	cmd := fmt.Sprintf(`helm uninstall %s --namespace %s`, chart.ReleaseName, chart.NameSpace)
	if err := execute(cmd); err != nil {
		return err
	}
	return nil
}

// addHelmRepo 添加 Helm 仓库
func addHelmRepo(chart *ChartArgs) error {
	// 替换为实际的仓库 URL
	cmd := fmt.Sprintf(`helm repo add %s %s`, chart.RepoName, chart.RepoURL)
	if err := execute(cmd); err != nil {
		return err
	}

	if err := execute(`helm repo update`); err != nil {
		return err
	}

	return nil
}

// installNewChart 使用 Helm 安装新的 Chart
func installNewChart(chart *ChartArgs, isExisted bool) error {
	var cmd string

	if isExisted {
		fileName := fmt.Sprintf("/tmp/%s.yaml", chart.ReleaseName)
		file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
		defer func() {
			file.Close()
			os.Remove(fileName)
		}()

		if _, err := file.Write([]byte(chart.Values)); err != nil {
			return err
		}

		// 执行 Helm install 命令
		cmd = fmt.Sprintf(`helm upgrade --install -n %s --create-namespace %s %s --version %s -f %s`, chart.NameSpace, chart.ReleaseName, chart.ChartName, chart.ChartVersion, fileName)
	} else {
		cmd = fmt.Sprintf(`helm upgrade --install -n %s --create-namespace %s %s --version %s`, chart.NameSpace, chart.ReleaseName, chart.ChartName, chart.ChartVersion)
	}

	// 将 values 写入临时文件并执行 安装或升级
	if err := execute(cmd); err != nil {
		return fmt.Errorf("helm install failed: %s", err)
	}
	return nil
}

// mergeChartValues  合并 values
func mergeChartValues(chart *ChartArgs) (err error) {
	var newValues, oldValues ChartValues

	if oldValues, err = getInstalledChartValues(chart.ReleaseName, chart.NameSpace); err != nil {
		return err
	}
	if newValues, err = getWillUpChartValues(chart.ChartName, chart.NameSpace); err != nil {
		return err
	}

	// 修改 values，将 image 字段替换为新的 image
	modifiedValues, err := modifyValues(oldValues, newValues)

	if err != nil {
		return err
	}
	chart.Values = modifiedValues
	return nil
}

// modifyValues 修改原有的 values，将 image 字段替换为新的 image
func modifyValues(oldValues, newValues ChartValues) (string, error) {
	// 创建新的 values
	resultValues := oldValues
	newImage := newValues["config"]["image"]
	resultValues["config"]["image"] = newImage
	// 转换为 YAML 格式
	modifiedValues, err := yaml.Marshal(resultValues)
	if err != nil {
		return "", err
	}

	return string(modifiedValues), nil
}

// getInstalledChartValues 获取创建是安装的 Helm release 的 values
func getInstalledChartValues(releaseName, namespace string) (ChartValues, error) {
	var out bytes.Buffer
	cmd := exec.Command("helm", "get", "values", releaseName, "-n", namespace, "-o", "yaml")
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	var values ChartValues
	if err := yaml.Unmarshal(out.Bytes(), &values); err != nil {
		return nil, err
	}
	return values, nil
}

// getWillUpChartValues 获取将要升级的 Helm release 的 values
func getWillUpChartValues(chartName, namespace string) (ChartValues, error) {
	var out bytes.Buffer
	cmd := exec.Command("helm", "show", "values", chartName, "-n", namespace)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	var values ChartValues
	if err := yaml.Unmarshal(out.Bytes(), &values); err != nil {
		return nil, err
	}
	return values, nil
}

func releaseExists(releaseName, namespace string) (bool, error) {
	cmd := exec.Command("helm", "list", "-n", namespace)
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}

	releases := strings.Split(string(output), "\n")
	for _, release := range releases {
		if strings.Contains(release, releaseName) {
			return true, nil
		}
	}

	return false, nil
}

func execute(command string) error {
	logger.Default.Info(context.Background(), "execute is: %s", command)

	var outBuffer, errBuffer bytes.Buffer
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer
	err := cmd.Run()

	logger.Default.Info(context.Background(), "execute [%s]: %s", command, outBuffer.String())
	if err != nil {
		logger.Default.Error(context.Background(), "execute [%s]: %s", command, errBuffer.String())
	}

	return err
}
