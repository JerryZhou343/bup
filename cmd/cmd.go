package cmd

import (
	"github.com/mfslog/prototool/internal/app"
	"github.com/spf13/cobra"
	"log"
)

func Run() {
	app, err := app.InitApp()
	if err != nil {
		log.Fatal("init App failed", err)
	}
	var (
		RootCmd = &cobra.Command{
			Use:        "prototool",
			ArgAliases: []string{"generate,fmt, config"},
			Args:       cobra.OnlyValidArgs,
		}

		generateCmd = &cobra.Command{
			Use:  "generate",
			Args: cobra.NoArgs,
			Run: func(cmd *cobra.Command, args []string) {
				app.Gen()
			},
		}

		fmtCmd = &cobra.Command{
			Use:  "fmt",
			Args: cobra.NoArgs,
			Run: func(cmd *cobra.Command, args []string) {
				app.Format()
			},
		}
	)

	RootCmd.AddCommand(generateCmd)
	RootCmd.AddCommand(fmtCmd)
	RootCmd.Execute()
}
