package syntax

import (
	"encoding/json"
	"fmt"
	"os"
)

type ChartArgs struct {
	ReleaseName  string
	ChartName    string
	ChartVersion string
	NameSpace    string
	RegistryURL  string
	Values       string
	RepoName     string
	RepoURL      string
	IsPublic     bool
}
type ChartValues map[string]map[string]interface{}

func UpdateAddonChart() (err error) {
	oldValues := ChartValues{
		"config": {"version": "v1.0.0"},
	}

	// 包装参数
	// var chart ChartArgs
	chart := new(ChartArgs)
	chart.initChartValues()
	ok, err := chart.isLogined()
	fmt.Printf("ok: %v,err: %v\n", ok, err)

	resultValues := oldValues
	if err = chart.joinURL(&resultValues, "config", "image", "value"); err != nil {
		return err
	}
	fmt.Printf("resultValues: %v\n", resultValues)
	return nil
}
func (chart *ChartArgs) initChartValues() error {
	chart.ChartName = "chart"
	chart.ChartVersion = "0.1.0"
	chart.NameSpace = "default"
	chart.Values = ``
	chart.RepoName = "my-repo"
	chart.RegistryURL = "192.168.2.244:8080"
	return nil
}
func (chart *ChartArgs) joinURL(resultValues *ChartValues, configName, valueName string, value interface{}) error {
	fmt.Printf("chart:%v \n", *chart)
	if img, ok := value.(string); ok {
		newURL := img
		// 解引用
		(*resultValues)[configName][valueName] = newURL
		return nil
	}
	return fmt.Errorf("JoinURL invalid value %v, should be a string", value)
}

type AuthConfig struct {
	Auths map[string]struct {
		Auth string `json:"auth"`
	} `json:"auths"`
}

func (chart *ChartArgs) isLogined() (bool, error) {
	configFilePath := fmt.Sprintf("%s/config.json", os.Getenv("HOME"))
	var (
		file []byte
		data AuthConfig
		err  error
	)

	if _, err = os.Stat(configFilePath); os.IsNotExist(err) {
		return false, err
	}

	if file, err = os.ReadFile(configFilePath); err != nil {
		return false, err
	}
	if err = json.Unmarshal(file, &data); err != nil {
		return false, err
	}
	fmt.Printf("data: %v\n", data)

	if _, exists := data.Auths[chart.RegistryURL]; exists {
		return true, nil
	} else {
		return false, nil
	}
}
