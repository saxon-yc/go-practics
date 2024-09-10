package syntax

import (
	"fmt"
)

type ChartArgs struct {
	ReleaseName  string
	ChartName    string
	ChartVersion string
	NameSpace    string
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

	chart := new(ChartArgs)
	// 包装参数
	resultValues := oldValues
	if err = chart.joinURL(&resultValues, "config", "image", "value"); err != nil {
		return err
	}
	fmt.Printf("resultValues: %v\n", resultValues)
	return nil
}
func (chart *ChartArgs) joinURL(resultValues *ChartValues, configName, valueName string, value interface{}) error {
	if img, ok := value.(string); ok {
		newURL := img
		// 解引用
		(*resultValues)[configName][valueName] = newURL
		return nil
	}
	return fmt.Errorf("JoinURL invalid value %v, should be a string", value)
}
