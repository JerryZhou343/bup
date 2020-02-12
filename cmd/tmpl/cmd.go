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
	tmpl = `#导入路径可以设置环境变量: $GOPATH 或者 ${GOPATH}
#导入目录
includes:
  - $GOPATH/src/github.com/bilibili/kratos/third_party
  - $IDLBASE/idl

#当前项目依赖的proto文件
depends:
  - xxxx/xxxx/xxxx.proto


excludes:
  - github.com/gogo/protobuf/gogoproto/gogo.proto
  - google

#编译设置
generate:
    go_options:
      extra_modifiers:
    plugins:
      - name: gofast
        flags: plugins=grpc
        output: ./api
`
)
