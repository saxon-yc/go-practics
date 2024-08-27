package cmd

import (
	"fmt"
	"log"

	"github.com/lithammer/dedent"
	"github.com/spf13/cobra"

	dbsvc "go-practics/internal/db"
	"go-practics/internal/model"
)

/* //go:embed components.yaml */
// var configFile embed.FS
/*
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
} */

func queryComponents(args []string) (err error) {
	var result []model.ClusterComponentFields
	if len(args) == 2 {
		result, err = dbsvc.NewDbServer().FindQkeComponents(args[0], args[1])
	} else {
		result, err = dbsvc.NewDbServer().FindQkeComponents("", "")
	}

	fmt.Printf("queryComponents:%+v\n", result)

	if err != nil {
		return err
	}
	return nil

}

func QueryComponent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query components",
		Short: "query: Import components configuration",
		Long: dedent.Dedent(`
		`),
		SilenceErrors: true,
		SilenceUsage:  true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("args: %v\n", args)
			if err := queryComponents(args); err != nil {
				log.Fatalf("Error adding components: %s", err)
			} else {
				log.Fatalln("Added components successfully")
			}
		},
	}
	return cmd
}
