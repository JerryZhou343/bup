package app

import (
	"github.com/mfslog/prototool/internal/compile"
	"github.com/mfslog/prototool/internal/conf"
	"github.com/mfslog/prototool/internal/format"
	"github.com/mfslog/prototool/internal/lint"
	"github.com/mfslog/prototool/internal/proto"
	"log"
	"os"
	"path/filepath"
)

type App struct {
	config    *conf.Config
	compiler  *compile.Compiler
	formatter *format.Formatter
}

func NewApp(config *conf.Config, compiler *compile.Compiler, formatter *format.Formatter) (*App, error) {
	return &App{
		config:    config,
		compiler:  compiler,
		formatter: formatter,
	}, nil
}

func (a *App) Format() {
	for _, itr := range a.config.Protos {
		absPath := filepath.Join(a.config.ImportPath, itr)
		_, err := os.Open(absPath)
		if err != nil {
			log.Println("can't access file", absPath)
			continue
		}
		a.formatter.FormatSignFile(absPath)
	}

}

func (a *App) Gen() {
	includePath := []string{}
	includePath = append(includePath, a.config.ImportPath)
	includePath = append(includePath, a.config.Includes...)
	descSource, err := proto.DescriptorSourceFromProtoFiles(includePath, a.config.Protos...)
	if err != nil {
		log.Fatalf("Failed to process proto source files. %v", err)
	}

	err = a.compiler.Compile(descSource)
	if err != nil {
		log.Fatalf("compile error %v", err)
	}

	return
}

func (a *App) Lint() {
	allAbsFile := []string{}
	for _, itr := range a.config.Protos {
		absPath := filepath.Join(a.config.ImportPath, itr)
		_, err := os.Open(absPath)
		if err != nil {
			log.Println("can't access file", absPath)
			continue
		}
		allAbsFile = append(allAbsFile, absPath)
	}
	text, err := lint.NewRunner().Run(allAbsFile, a.config.LintCfg)
	if err != nil {
		log.Fatalln(err)
	}

	for _, itr := range text {
		log.Println(itr)
	}
}
