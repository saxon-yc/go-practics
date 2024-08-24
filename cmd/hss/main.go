package main

import (
	"fmt"
	"go-practics/cmd/hss/app"
	"os"
)

func main() {
	rootCmd := app.NewHssCommand()

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
