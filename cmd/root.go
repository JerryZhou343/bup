package cmd

import (
	"github.com/mfslog/prototool/internal/app"
	"github.com/spf13/cobra"
)

var (
	App *app.App
)

var (
	RootCmd = &cobra.Command{
		Use:  "prototool",
		Args: cobra.NoArgs,
	}
)

func init() {
	App, _ = app.InitApp()
}

func Run() {
	RootCmd.AddCommand(generateCmd)
	RootCmd.AddCommand(fmtCmd)
	RootCmd.AddCommand(lintCmd)
	RootCmd.AddCommand(configCmd)
	RootCmd.Execute()
}
