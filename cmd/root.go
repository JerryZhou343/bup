package cmd

import (
	"fmt"
	"github.com/JerryZhou343/prototool/internal/app"
	"github.com/spf13/cobra"
)

var (
	App   *app.App
	major = 2
	minor = 1
	patch = 0
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

	RootCmd.Version = fmt.Sprintf("v%d.%d.%d", major, minor, patch)
	RootCmd.Execute()
}
