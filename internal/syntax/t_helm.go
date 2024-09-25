package syntax

import (
	"fmt"
	"go-practics/internal/util"
	"os"
	"os/exec"
	"strings"
)

const helmPath = "/opt/homebrew/bin/helm"

func NewChartCmd() {
	checkHelmExist()
	checkHelmVersion()
}

func checkHelmExist() {
	if _, err := os.Stat(helmPath); os.IsNotExist(err) {
		fmt.Println("Helm not found, please install Helm first.")
		os.Exit(1)
	}
}

type HelmVersion struct {
	Client ClientInfo `json:"client"`
}

type ClientInfo struct {
	Version      string `json:"version"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	GoVersion    string `json:"goVersion"`
}

func checkHelmVersion() error {
	cmd := exec.Command(helmPath, "version", "--short")
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	helmVersion := string(output)
	version := strings.Split(helmVersion, "+")
	fmt.Printf("version[0]: %v\n", version[0])
	// for _, info := range helmInfos {
	// 	if strings.Contains(info, chart.ReleaseName) {
	// 		return nil
	// 	}
	// }
	if ok, err := util.CompareVersionStr(version[0], "v3.18.0"); err != nil {
		return err
	} else {
		if !ok {
			tip := fmt.Sprintf("The current Helm version is %s. Please upgrade the helm version to a version greater than v3.8.0", version[0])
			fmt.Println(tip)
			os.Exit(1)
		}
	}
	return nil

}
