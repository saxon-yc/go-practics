package external

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

/* viper包，主要用于读/写文件 */
func NewViper(proxy string) {
	// 读取配置
	viper.SetConfigFile(proxy)
	viper.AddConfigPath(".")
	if error := viper.ReadInConfig(); error != nil {
		fmt.Printf("read config file %s failed: %v\n", viper.ConfigFileUsed(), error)
		os.Exit(1)
	}
	fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())
}

func MyViper() {
	fmt.Printf("viper.GetInt(\"app.port\"): %v\n", viper.GetInt("app.port")) // 9999

	fmt.Printf("viper.GetString(\"app.pre\"): %v\n", viper.GetString("app.pre"))
}
