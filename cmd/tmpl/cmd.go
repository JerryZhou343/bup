package tmpl

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	InitCmd = &cobra.Command{
		Use:   "init",
		Short: "create new idl.yaml",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func run() {
	var (
		err  error
		file *os.File
	)
	//file, err = os.OpenFile("./idl.yaml", os.O_CREATE, os.)
	file, err = os.Create("idl.yaml")
	if err != nil {
		log.Fatalf("create file failed: %v", err)
	}

	_, err = file.WriteString(tmpl)
	if err != nil {
		log.Fatalf("write file failed: %v", err)
	}
}

var (
	tmpl = `#idl 基础导入路径,路径可以设置环境变量: $GOPATH 或者 ${GOPATH}
import_path: /Users/apple/workbench/src/github.com/mfslog/prototool/example/idl
#当前项目依赖的proto文件
depends:
  - usermgt/passport/passport.proto
#依赖第三方目录
includes:

#编译设置
generate:
    go_options:
      extra_modifiers:
        usermgt/user.proto: code.hyfco.com/usermgt/output/usermgt/userbase
        usermgt/code.proto: code.hyfco.com/usermgt/output/usermgt/userbase
    plugins:
      - name: gofast
        flags: plugins=grpc
        output: ./output
`
)
