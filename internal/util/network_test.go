package util

import (
	"fmt"
	"strings"
	"testing"
)

func TestPingHost(t *testing.T) {
	url := "www.baidu.com" // 替换为你要测试的主机或 IP 地址

	// 从 URL 中提取主机名
	host := url // 如果是 IP 地址或域名直接使用
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		host = strings.Split(url, "//")[1]
		host = strings.Split(host, "/")[0] // 取出主机名部分
	}

	fmt.Printf("Pinging host %s...\n", host)
	isReachable, err := PingHost(host)
	if err != nil {
		fmt.Printf("Error is: %s\n", err)
	} else if isReachable {
		fmt.Println("Host is reachable!")
	} else {
		fmt.Println("Host is not reachable.")
	}
}

func TestCheckUrl(t *testing.T) {
	url := "http://example.com" // 替换为你要测试的 URL

	fmt.Printf("Checking URL %s...\n", url)
	isReachable, err := CheckURL(url)
	if err != nil {
		fmt.Printf("Error is: %s\n", err)
	} else if isReachable {
		fmt.Println("URL is reachable!")
	} else {
		fmt.Println("URL is not reachable.")
	}
}

func TestCheckImageExists(t *testing.T) {
	image := "nginx" // 要检查的 Nginx 镜像

	exists, err := CheckImageExists(image)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	if exists {
		fmt.Printf("The image '%s' exists on Docker Hub.\n", image)
	} else {
		fmt.Printf("The image '%s' does not exist on Docker Hub.\n", image)
	}
}
