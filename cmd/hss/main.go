package main

import (
	chart "go-practics/cmd/hss/push"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	// rootCmd := app.NewHssCommand()

	// if err := rootCmd.Execute(); err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	os.Exit(1)
	// }
	cmd := NewInitCmd()
	err := cmd.Execute()
	if err != nil {
		os.Exit(-1)
	}

}

func NewInitCmd() *cobra.Command {
	root := &cobra.Command{
		Use:   "qke",
		Short: "Init QKE ",
		Example: `
`,
		Version: "v0.0.1",
	}

	root.AddCommand(chart.NewChartCmd())

	return root
}
