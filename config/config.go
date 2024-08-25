// 配置文件
package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var ConfigFile ConfigModel

// 一般把读写初始化配置 写在config文件夹下
func New(proxyPath, comptPath string) {
	// 读取配置
	viper.SetConfigFile(proxyPath)
	viper.AddConfigPath(".")
	if error := viper.ReadInConfig(); error != nil {
		log.Fatalf("read config file %s; failed: %v", viper.ConfigFileUsed(), error)
		os.Exit(1)
	}

	viper.SetConfigFile(comptPath)
	viper.AddConfigPath(".")
	if err := viper.MergeInConfig(); err != nil {
		log.Fatalf("merge config file[%s], failed[%s]\n", comptPath, err)
		os.Exit(1)
	}
	// fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())

	if err := viper.Unmarshal(&ConfigFile); err != nil {
		log.Fatalf("unmarshal config error: %v\n", err)
		os.Exit(1)
	}
	// fmt.Printf("viper.GetString(\"app.port\"): %v\n", viper.GetString("app.port"))
}
