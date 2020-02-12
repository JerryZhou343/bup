package generate

import (
	"fmt"
	"github.com/mfslog/prototool/pkg/conf"
	"github.com/mfslog/prototool/pkg/proto"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

type Executor struct {
	includes       []string
	protoDesc      proto.DescriptorSource
	extraModifiers map[string]string
	plugins        []conf.Plugin
	excludes       map[string]struct{}
}

type metaCmd struct {
	file           string
	args           []string
	extraModifiers string
}

func (e *Executor) generate() {
	var (
		err error
	)
	cmds := e.generateCmd()
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
		for _, cmd := range cmds {
			if len(cmd.extraModifiers) > 0 {
				arg = arg + "," + cmd.extraModifiers + ":" + outputPath
			} else {
				arg = arg + ":" + outputPath
			}
			//输出路径
			tmp := append(cmd.args, arg)
			tmp = append(tmp, cmd.file)
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

func (e *Executor) skip(name string) bool {
	//完整匹配
	if _, ok := e.excludes[name]; ok {
		return true
	}
	dir := path.Dir(name)
	for {
		//log.Printf("dir %s", dir)
		if _, ok := e.excludes[dir]; ok {
			return true
		}
		dir = path.Dir(dir)
		if dir == string(os.PathSeparator) || dir == "." || dir == "" {
			return false
		}
	}
}

func (e *Executor) generateCmd() []*metaCmd {
	var (
		M   string
		ret []*metaCmd
	)
	for name, fileDesc := range e.protoDesc.Files() {
		//判断过滤
		if e.skip(name) {
			continue
		}
		//生成命令
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
		for _, itr := range e.includes {
			args = append(args, "-I"+itr)
		}

		cmd := metaCmd{
			file:           name,
			args:           args,
			extraModifiers: M,
		}
		ret = append(ret, &cmd)
	}
	return ret
}
