package generate

import (
	"github.com/mfslog/prototool/pkg/conf"
	"github.com/mfslog/prototool/pkg/proto"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	GenCmd = &cobra.Command{
		Use:     "generate",
		Aliases: []string{"g"},
		Short:   "compile proto file",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func run() {
	err := conf.Init()
	if err != nil {
		log.Fatalf("init  read idl.yaml failed [%v]", err)
	}
	executor := new(Executor)
	executor.excludes = make(map[string]struct{})
	allFilePaths := []string{}
	executor.extraModifiers = conf.IDLConfig.Generate.GoOptions.ExtraModifiers
	executor.plugins = conf.IDLConfig.Generate.Plugins
	for _, itr := range conf.IDLConfig.Excludes {
		executor.excludes[itr] = struct{}{}
	}
	for _, itr := range conf.IDLConfig.Includes {
		executor.includes = append(executor.includes, os.ExpandEnv(itr))
		allFilePaths = append(allFilePaths, os.ExpandEnv(itr))
	}

	descSource, err := proto.DescriptorSourceFromProtoFiles(allFilePaths, conf.IDLConfig.Depends...)
	if err != nil {
		log.Println(err, "Failed to process proto source files.")
		os.Exit(1)
	}
	executor.protoDesc = descSource
	executor.generate()
}
