package util

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// pingHost 使用 ping 命令检查目标主机的连通性
func PingHost(host string) (bool, error) {
	if host == "" {
		return false, errors.New("host cannot be empty")
	}
	var cmd *exec.Cmd
	// 执行 ping 命令，-c 选项指定发送的包数（Linux/macOS），-n 选项用于 Windows
	switch strings.ToLower(runtime.GOOS) {
	case "linux", "darwin":
		cmd = exec.Command("ping", "-c", "4", host)
	case "windows":
		cmd = exec.Command("ping", "-n", "4", host)
	default:
		cmd = exec.Command("ping", "-c", "4", host)
	}

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		errStr := err.Error()
		if strings.Contains(errStr, "exit status 2") {
			errStr = "please provide a currently host"
		} else if strings.Contains(errStr, "exit status 68") {
			errStr = "the network is broken"
		}
		return false, errors.New(errStr)
	}

	// 检查输出中是否包含 "0% packet loss"
	output := out.String()
	fmt.Printf("output: %v\n", output)
	if strings.Contains(output, "0% packet loss") {
		return true, nil
	}

	return false, nil
}

// checkURL 使用 HTTP GET 请求检查 URL 的连通性
func CheckURL(url string) (bool, error) {
	// 发送 HTTP GET 请求
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)

	if err != nil {
		return false, err
	}
	defer resp.Body.Close() // 确保在函数结束时关闭响应体
	// 检查 HTTP 状态码
	if resp.StatusCode == http.StatusOK {
		return true, nil
	}
	return false, nil
}

type DockerImage struct {
	Name string `json:"name"`
}

func CheckImageExists(image string) (bool, error) {
	url := fmt.Sprintf("http://xxx/%s", image)
	resp, err := http.Get(url)
	if err != nil {
		return false, fmt.Errorf("failed to check image: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, nil // 镜像存在
	} else if resp.StatusCode == http.StatusNotFound {
		return false, nil // 镜像不存在
	}

	return false, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
}
