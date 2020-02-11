package generate

import (
	"fmt"
	"github.com/mfslog/prototool/pkg/conf"
	"github.com/mfslog/prototool/pkg/proto"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Executor struct {
	importPath     string
	includes       []string
	protoDesc      proto.DescriptorSource
	extraModifiers map[string]string
	plugins        []conf.Plugin
}

func (e *Executor) generate() {
	var (
		M   string
		err error
	)
	for name, fileDesc := range e.protoDesc.Files() {
		//destFilePath := filepath.Join(e.importPath, name)
		destFilePath :=  name
		fs := fileDesc.GetDependencies()
		var ms []string
		for _, fd := range fs {
			dependName := fd.GetName()
			opt := fd.GetFileOptions().GoPackage
			if opt == nil {
				os.Exit(1)
			}
			if m, ok := e.extraModifiers[dependName]; ok {
				ms = append(ms, "M"+dependName+"="+m)
			}
		}
		if len(ms) > 0 {
			M = strings.Join(ms, ",")
		}

		var args []string
		//-I 包含路径
		args = []string{"-I" + e.importPath}
		for _, itr := range e.includes {
			args = append(args, "-I"+itr)
		}

		for _, itr := range e.plugins {
			//插件
			arg := fmt.Sprintf("--%s_out=", itr.Name)
			outputPath := os.ExpandEnv(itr.Output)
			err = os.MkdirAll(outputPath, os.ModePerm)
			if err != nil {
				log.Fatalf("MkdirAll failed [%v]", err)
				break
			}
			//参数
			arg = arg + itr.Flags
			if len(M) > 0 {
				arg = arg + "," + M + ":" + outputPath
			} else {
				arg = arg + ":" + outputPath
			}
			//输出路径
			tmp := append(args, arg)
			tmp = append(tmp, destFilePath)
			cmd := exec.Command("protoc", tmp...)
			log.Println(cmd.Args)
			out, err := cmd.CombinedOutput()
			if err != nil {
				log.Println("Error:", err)
				log.Println(string(out))
			}
		}

	}
}
