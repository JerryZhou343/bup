package cmd

import (
	"github.com/JerryZhou343/prototool/internal/flags"
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

func init() {
	generateCmd.PersistentFlags().StringSliceVarP(&flags.SrcFiles, "files", "f",
		[]string{}, "相对于导入目录下的源文件")
	generateCmd.PersistentFlags().StringSliceVarP(&flags.SrcDirectories, "directories", "d",
		[]string{}, "相对于导入目录下源文件目录")
}
