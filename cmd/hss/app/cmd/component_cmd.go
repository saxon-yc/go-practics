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
	var config model.Components
	if err := readFile(); err != nil {
		return err
	} else {
		if err := viper.Unmarshal(&config); err != nil {
			log.Fatalf("Error unmarshalling config: %s", err)
			return err
		}

		// fmt.Printf("Components: %v\n", config.Components)
		if err := dbsvc.NewDbServer().AddComponents(config.Components); err != nil {
			log.Fatalf("Error adding components: %s", err)
			return err
		}
		return nil
	}

}

func ComponentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import",
		Short: "import: Import components configuration",
		Long: dedent.Dedent(`
		`),
		SilenceErrors: true,
		SilenceUsage:  true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("importComponents(): %v\n", importComponents())
		},
	}
	return cmd
}
