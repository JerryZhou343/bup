//go:build wireinject
// +build wireinject

package app

import (
	"github.com/JerryZhou343/bup/internal/compile"
	"github.com/JerryZhou343/bup/internal/conf"
	"github.com/JerryZhou343/bup/internal/format"
	"github.com/google/wire"
)

func InitApp() (*App, error) {
	panic(wire.Build(conf.NewConfig, compile.NewCompiler, format.NewFormatter, NewApp))
}
