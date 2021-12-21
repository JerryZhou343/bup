package compile

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/JerryZhou343/prototool/internal/conf"
	"github.com/JerryZhou343/prototool/internal/proto"
	"github.com/JerryZhou343/prototool/internal/wkt"
	"github.com/jhump/protoreflect/desc"
)

type Compiler struct {
	config *conf.Config
}

type metaCmd struct {
	file           string
	args           []string
	extraModifiers string
}

func NewCompiler(cfg *conf.Config) *Compiler {
	ret := &Compiler{
		config: cfg,
	}

	return ret
}

func (c *Compiler) Compile(desc proto.DescriptorSource, deleteDirectory bool) (err error) {

	outputPath := os.ExpandEnv(c.config.Generate.Output)

	if deleteDirectory {
		err = os.RemoveAll(outputPath)
		if err != nil {
			log.Fatalf("remove path failed. [%s] ", outputPath)
		}
	}

	err = os.MkdirAll(outputPath, os.ModePerm)
	if err != nil {
		log.Fatalf("MkdirAll failed [%v]", err)
	}

	log.Println("out path ", outputPath)
	for _, itr := range c.config.Generate.Plugins {
		metaCmds := c.generateCmd(desc, itr.Type)
		//插件
		arg := fmt.Sprintf("--%s_out=", itr.Name)

		//参数
		arg = arg + itr.Flags
		var optArg string
		for _, cmd := range metaCmds {
			optArg = ""
			if len(cmd.extraModifiers) > 0 {
				optArg = arg + "," + cmd.extraModifiers + ":" + outputPath
			} else {
				optArg = arg + ":" + outputPath
			}
			//输出路径
			tmp := append(cmd.args, optArg)
			tmp = append(tmp, cmd.file)
			tmpCmd := exec.Command("protoc", tmp...)
			log.Println(tmpCmd.String())
			out, err := tmpCmd.CombinedOutput()
			if err != nil {
				log.Println("compile Error:", err)
				log.Println(string(out))
			}
		}
	}
	return nil
}

func (e *Compiler) generateCmd(descriptor proto.DescriptorSource, typ string) []*metaCmd {
	var (
		M           string
		ret         []*metaCmd
		prjFiles    map[string]struct{}
		prjFileDesc map[string]*desc.FileDescriptor
	)

	prjFiles = make(map[string]struct{})
	prjFileDesc = make(map[string]*desc.FileDescriptor)

	for _, name := range e.config.Protos {
		prjFiles[name] = struct{}{}
	}

	for name, fileDesc := range descriptor.Files() {
		prjFileDesc[name] = fileDesc
	}

	for name, fileDesc := range descriptor.Files() {
		M = ""
		if _, ok := wkt.Filenames[name]; ok {
			continue
		}

		if _, ok := e.config.IngoreMap[name]; ok {
			continue
		}

		//生成命令
		fs := fileDesc.GetDependencies()
		var ms []string
		for _, fd := range fs {
			//和当前编译文件在同一个包下面，跳过
			if fd.GetFileOptions().GetGoPackage() ==
				fileDesc.GetFileOptions().GetGoPackage() {
				continue
			}

			dependName := fd.GetName()
			//指定依赖
			if m, ok := e.config.Generate.GoOptions.ExtraModifiers[dependName]; ok {
				ms = append(ms, "M"+dependName+"="+m)
				continue
			}
			//wkt
			if typ == "go" {
				if m, ok := wkt.FilenameToGoModifierMap[dependName]; ok {
					ms = append(ms, "M"+dependName+"="+m)
				}
			} else {
				if m, ok := wkt.FilenameToGogoModifierMap[dependName]; ok {
					ms = append(ms, "M"+dependName+"="+m)
				}
			}
			//当前声明自身依赖
			if _, ok := prjFiles[dependName]; ok {
				//正确生成描述
				if dfd, okd := prjFileDesc[dependName]; okd {
					//生成依赖包名
					tmp := path.Join(e.config.GoModule, e.config.Generate.Output,
						*dfd.GetFileOptions().GoPackage)
					ms = append(ms, "M"+dependName+"="+tmp)
				}
			}
		}
		if len(ms) > 0 {
			M = strings.Join(ms, ",")
		}

		var args []string
		//-I 包含路径
		for _, itr := range e.config.Includes {
			args = append(args, "-I"+itr)
		}
		cmd := metaCmd{
			file:           filepath.Join(e.config.ImportPath, name),
			args:           args,
			extraModifiers: M,
		}
		ret = append(ret, &cmd)
	}
	return ret
}
