package conf

var (
	tmpl = `#导入路径可以设置环境变量: $GOPATH 或者 ${GOPATH}
#项目基础导入目录
import_path:$IDLBASE/idl

#当前项目依赖的proto文件
protos:
  - xxxx/xxxx/xxxx.proto

#依赖导入目录
includes:
  - $GOPATH/src/github.com/bilibili/kratos/third_party



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
