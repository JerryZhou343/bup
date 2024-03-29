# BuildYourProtocol
proto 辅助工具，根据配置文件的依赖描述，编译，格式化proto，还可以对proto进行拼写检查

# 安装
```shell script
go install  github.com/JerryZhou343/bup@latest
```

# 命令
请查看help

```shell script
bup -h
Usage:
  bup [command]

Available Commands:
  config      生成配置文件
  fmt         格式化proto
  gen    编译proto
  help        Help about any command
  lint        对proto进行拼写检查

Flags:
  -h, --help   help for bup

Use "bp [command] --help" for more information about a command.
```

# 依赖

格式化功能依赖clang-format;请下在clang-format 后将clang-format 放到PATH 环境变量下；

# 参考
###  同类工具

[uber bup](https://github.com/uber/bup)
[buf](https://github.com/bufbuild/buf)

### 拼写检查手册

[Google API Linter](https://linter.aip.dev/)

### llvm 

[clang-format 下载](https://github.com/llvm/llvm-project)

[clang-format 说明](https://clang.llvm.org/docs/ClangFormat.html)
