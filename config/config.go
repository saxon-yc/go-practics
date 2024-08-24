// 配置文件
package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// 一般把读写初始化配置 写在config文件夹下
func New(path string) {
	// 读取配置
	viper.SetConfigFile(path)
	viper.AddConfigPath(".")
	if error := viper.ReadInConfig(); error != nil {
		fmt.Printf("read config file %s; failed: %v", viper.ConfigFileUsed(), error)
		os.Exit(1)
	}
	fmt.Printf("Using config file: %s", viper.ConfigFileUsed())
	// fmt.Printf("viper.GetString(\"app.port\"): %v\n", viper.GetString("app.port"))
}
