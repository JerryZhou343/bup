package conf

var (
	tmpl = `###################基础配置#####################
#导入路径可以设置环境变量: $GOPATH 或者 ${GOPATH}
#项目基础导入目录
import_path: $IDL_PATH

#当前项目依赖的proto文件
protos:
  - usermgt/passport/passport.proto
  - usermgt/code.proto
  - usermgt/user.proto
#依赖导入目录
includes:
  - $GOPATH/src/github.com/bilibili/kratos/third_party


####################lint配置####################
lint:
# 检查规则
  rules:
    enable:

    disable:


####################编译配置####################
generate:
    go_options:
      extra_modifiers:
    plugins:
      - name: gofast
        flags: plugins=grpc
        output: ./genproto
`
)
