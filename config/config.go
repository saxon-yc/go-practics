// 配置文件
package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func New(path string) {
	// 读取配置
	viper.SetConfigFile(path)
	viper.AddConfigPath(".")
	if error := viper.ReadInConfig(); error != nil {
		fmt.Printf("read config file %s failed: %v", viper.ConfigFileUsed(), error)
		os.Exit(1)
	}
	fmt.Printf("Using config file: %s", viper.ConfigFileUsed())
	// fmt.Printf("viper.GetString(\"app.port\"): %v\n", viper.GetString("app.port"))
}
