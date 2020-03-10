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
			Use: "prototool",
		}

		generateCmd = &cobra.Command{
			Use:   "generate",
			Short: "编译proto",
			Args:  cobra.NoArgs,
			Run: func(cmd *cobra.Command, args []string) {
				app.Gen()
			},
		}

		fmtCmd = &cobra.Command{
			Use:   "fmt",
			Short: "格式化proto",
			Args:  cobra.NoArgs,
			Run: func(cmd *cobra.Command, args []string) {
				app.Format()
			},
		}

		lintCmd = &cobra.Command{
			Use:  "lint",
			Args: cobra.NoArgs,
			Run: func(cmd *cobra.Command, args []string) {
				err = app.Lint()
				if err != nil {
					log.Fatal(err)
				}
			},
		}

		configCmd = &cobra.Command{
			Use:  "config",
			Args: cobra.NoArgs,
			Run: func(cmd *cobra.Command, args []string) {
				app.Config()
			},
		}
	)

	RootCmd.AddCommand(generateCmd)
	RootCmd.AddCommand(fmtCmd)
	RootCmd.AddCommand(lintCmd)
	RootCmd.AddCommand(configCmd)
	RootCmd.Execute()
}
