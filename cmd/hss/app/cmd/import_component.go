package cmd

import (
	"bytes"
	"embed"
	"fmt"
	"log"

	"github.com/lithammer/dedent"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	dbsvc "go-practics/internal/db"
	"go-practics/internal/model"
)

//go:embed components.yaml
var configFile embed.FS

func readFile() error {
	// 读取嵌入的 YAML 文件
	byt, err := configFile.ReadFile("components.yaml")
	if err != nil {
		log.Fatalf("Error reading embedded config file: %s", err)
		return err
	}

	// 将数据写入到 Viper
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(bytes.NewReader(byt)); err != nil {
		log.Fatalf("Error reading config data: %s", err)
		return err
	}

	return nil
}

func importComponents() error {
	// var result config.ConfigModel
	// fmt.Printf("result:%+v\n", result.Components)
	// if err := dbsvc.NewDbServer().AddComponents(&configFile.Components); err != nil {
	// 	return err
	// }
	// return nil
	var config model.Components
	if err := readFile(); err != nil {
		return err
	} else {
		if err := viper.Unmarshal(&config); err != nil {
			return err
		}

		fmt.Printf("Components: %v\n", config.Components)
		if err := dbsvc.NewDbServer().AddComponents(config.Components); err != nil {
			return err
		}
		return nil
	}
}

func ImportComponent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import components",
		Short: "import: Import components configuration",
		Long: dedent.Dedent(`
		`),
		SilenceErrors: true,
		SilenceUsage:  true,
		Run: func(cmd *cobra.Command, args []string) {
			if err := importComponents(); err != nil {
				log.Fatalf("Error adding components: %s", err)
			} else {
				log.Fatalln("Added components successfully")
			}
		},
	}
	return cmd
}
