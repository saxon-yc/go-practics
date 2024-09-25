package chart

import (
	"github.com/spf13/cobra"
)

const helmPath = "/opt/homebrew/bin/helm"
const suppotOciVersuon = "v3.8.0"

func NewChartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "push",
		Aliases: []string{"p"},
		Short:   "push",
		Long:    "you can push chart to registry",
		Example: `qke push chart -p ./chart/ 192.168.0.1:8999 qke admin Harbor12345
# qke push chart -p chartPath registryURL projectName username password`,
	}
	cmd.AddCommand(NewPushChartCmd())

	return cmd
}
