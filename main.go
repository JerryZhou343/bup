package main

import (
	"github.com/mfslog/prototool/cmd/generate"
	"github.com/mfslog/prototool/cmd/tmpl"
	"github.com/spf13/cobra"
)

func main() {

	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(generate.GenCmd)
	rootCmd.AddCommand(tmpl.InitCmd)
	rootCmd.Execute()
}
