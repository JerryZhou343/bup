//+build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/JerryZhou343/prototool/internal/compile"
	"github.com/JerryZhou343/prototool/internal/conf"
	"github.com/JerryZhou343/prototool/internal/format"
)

func InitApp() (*App, error) {
	panic(wire.Build(conf.NewConfig, compile.NewCompiler, format.NewFormatter, NewApp))
}
