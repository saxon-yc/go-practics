package chart

import (
	"fmt"
	"go-practics/internal/util"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func NewPushChartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "chart",
		Aliases: []string{"c"},
		Short:   "chart",
		Long:    "you can push chart to registry",
		Example: `
			qke push chart -p ./chart/ 192.168.0.1:8999 qke admin zhu88jie!
			# qke push chart -p chartPath registryURL projectName username password
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := checkHelmExist(); err != nil {
				os.Exit(1)
			}
			if err := checkHelmVersion(); err != nil {
				os.Exit(1)
			}
			if err := checkChartExist(args); err != nil {
				os.Exit(1)
			}
			if err := loginRegistry(args); err != nil {
				os.Exit(1)
			}
			if err := pushChart(args); err != nil {
				os.Exit(1)
			}

			return nil
		},
	}
	return cmd
}

func checkHelmExist() error {
	if _, err := os.Stat(helmPath); os.IsNotExist(err) {
		fmt.Println("Helm not found, please install chart first.")
		return err
	}
	return nil
}

func checkHelmVersion() error {
	cmd := exec.Command(helmPath, "version", "--short")
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	versions := strings.Split(string(output), "+")
	helmVersion := versions[0]
	if ok, err := util.CompareVersionStr(helmVersion, suppotOciVersuon); err != nil {
		return err
	} else {
		if !ok {
			tip := fmt.Sprintf("The current Helm version is %s. Please upgrade the helm version to a version greater than %s", helmVersion, suppotOciVersuon)
			fmt.Println(tip)
			return err
		}
	}
	return nil
}

func checkChartExist(args []string) error {
	if _, err := os.Stat(args[0]); os.IsNotExist(err) {
		fmt.Println("Chart not found, please download chart first.")
		return err
	}
	return nil
}

func loginRegistry(args []string) error {
	registryURL := args[1]
	username := args[3]
	passwd := args[4]

	cmd := exec.Command("helm", "registry", "login", "-u", username, "-p", passwd, registryURL, "--insecure")

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func pushChart(args []string) (err error) {
	var files []fs.DirEntry
	chartDir := args[0]
	registryURL := args[1]
	projName := args[2]
	harborURL := fmt.Sprintf("oci://%s/%s", registryURL, projName)

	if files, err = os.ReadDir(chartDir); err != nil {
		return err
	}
	for _, file := range files {
		// 检查文件是否是 chart 包（以 .tgz 结尾）
		if strings.HasSuffix(file.Name(), ".tgz") {
			chartPath := filepath.Join(chartDir, file.Name())
			fmt.Printf("Pushing chart: %s\n", chartPath)

			// 使用 helm push 命令推送 chart 到 Harbor
			cmd := exec.Command("helm", "push", chartPath, harborURL, "--plain-http")

			output, err := cmd.CombinedOutput()
			if err != nil {
				log.Printf("Failed to push chart %s: %v\nOutput: %s", file.Name(), err, output)
				return err
			} else {
				fmt.Printf("Successfully pushed chart: %s\n", file.Name())
			}
		}
	}

	return nil
}
