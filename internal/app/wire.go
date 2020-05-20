//+build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/mfslog/prototool/internal/compile"
	"github.com/mfslog/prototool/internal/conf"
	"github.com/mfslog/prototool/internal/format"
)

func InitApp() (*App, error) {
	panic(wire.Build(conf.NewConfig, compile.NewCompiler, format.NewFormatter, NewApp))
}
