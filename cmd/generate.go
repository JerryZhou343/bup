package cmd

import (
	"github.com/spf13/cobra"
)

var (
	generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "编译proto",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			App.Gen()
		},
	}
)
