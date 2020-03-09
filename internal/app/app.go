package app

import (
	"fmt"
	"github.com/googleapis/api-linter/lint"
	"github.com/googleapis/api-linter/rules"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/mfslog/prototool/internal/compile"
	"github.com/mfslog/prototool/internal/conf"
	"github.com/mfslog/prototool/internal/format"
	"github.com/mfslog/prototool/internal/proto"
	"github.com/pkg/errors"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
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
	var (
		err error
	)
	err = a.config.Load()
	if err != nil {
		log.Fatal(err)
	}
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
	var (
		err error
	)
	err = a.config.Load()
	if err != nil {
		log.Fatal(err)
	}

	descSource, err := proto.DescriptorSourceFromProtoFiles(a.config.Includes, a.config.Protos...)
	if err != nil {
		log.Fatalf("Failed to process proto source files. %v", err)
	}

	err = a.compiler.Compile(descSource)
	if err != nil {
		log.Fatalf("compile error %v", err)
	}

	return
}

func (a *App) Lint() (err error) {
	err = a.config.Load()
	if err != nil {
		log.Fatal(err)
	}
	lintCfgs := lint.Configs{}
	// Add configs for the enabled rules.
	lintCfgs = append(lintCfgs, lint.Config{
		EnabledRules: a.config.Lint.Rules.Enable,
	})
	// Add configs for the disabled rules.
	lintCfgs = append(lintCfgs, lint.Config{
		DisabledRules: a.config.Lint.Rules.Disable,
	})
	// Prepare proto import lookup.
	fs, err := loadFileDescriptors()
	if err != nil {
		return err
	}
	lookupImport := func(name string) (*desc.FileDescriptor, error) {
		if f, found := fs[name]; found {
			return f, nil
		}
		return nil, fmt.Errorf("%q is not found", name)
	}

	var errorsWithPos []protoparse.ErrorWithPos
	var lock sync.Mutex
	// Parse proto files into `protoreflect` file descriptors.
	p := protoparse.Parser{
		ImportPaths:           a.config.Includes,
		IncludeSourceCodeInfo: true,
		LookupImport:          lookupImport,
		ErrorReporter: func(errorWithPos protoparse.ErrorWithPos) error {
			// Protoparse isn't concurrent right now but just to be safe for the future.
			lock.Lock()
			errorsWithPos = append(errorsWithPos, errorWithPos)
			lock.Unlock()
			// Continue parsing. The error returned will be protoparse.ErrInvalidSource.
			return nil
		},
	}

	// Resolve file absolute paths to relative ones.
	protoFiles, err := protoparse.ResolveFilenames(a.config.Includes, a.config.Protos...)
	if err != nil {
		return err
	}
	fd, err := p.ParseFiles(protoFiles...)
	if err != nil {
		if err == protoparse.ErrInvalidSource {
			if len(errorsWithPos) == 0 {
				return errors.New("got protoparse.ErrInvalidSource but no ErrorWithPos errors")
			}
			// TODO: There's multiple ways to deal with this but this prints all the errors at least
			errStrings := make([]string, len(errorsWithPos))
			for i, errorWithPos := range errorsWithPos {
				errStrings[i] = errorWithPos.Error()
			}
			return errors.New(strings.Join(errStrings, "\n"))
		}
		return err
	}

	// Create a linter to lint the file descriptors.
	globalRule := lint.NewRuleRegistry()
	rules.Add(globalRule)
	l := lint.New(globalRule, lintCfgs)
	results, err := l.LintProtos(fd...)
	if err != nil {
		return err
	}

	// Determine the output for writing the results.
	// Stdout is the default output.
	w := os.Stdout

	// Determine the format for printing the results.
	// YAML format is the default.
	marshal := getOutputFormatFunc("yaml")

	// Print the results.
	b, err := marshal(results)
	if err != nil {
		return err
	}
	if _, err = w.Write(b); err != nil {
		return err
	}
	return nil
}

func (a *App) Config() {
	err := a.config.Output()
	log.Fatal(err)
	return
}
