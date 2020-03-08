package conf

import (
	"github.com/mfslog/prototool/internal/lint"
	"github.com/mfslog/prototool/internal/strs"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Plugin struct {
	Name   string `yaml:"name"`
	Output string `yaml:"output"`
	Flags  string `yaml:"flags"`
}

type GoOptions struct {
	ExtraModifiers map[string]string `yaml:"extra_modifiers"`
}

type Generate struct {
	GoOptions GoOptions `yaml:"go_options"`
	Plugins   []Plugin  `yaml:"plugins"`
}

type Config struct {
	ImportPath string   `yaml:"import_path"`
	Protos     []string `yaml:"protos"`
	Includes   []string `yaml:"includes"`
	Generate   Generate `yaml:"generate"`
	Lint       struct {
		Group   string `json:"group,omitempty" yaml:"group,omitempty"`
		Ignores []struct {
			ID    string   `json:"id,omitempty" yaml:"id,omitempty"`
			Files []string `json:"files,omitempty" yaml:"files,omitempty"`
		} `json:"ignores,omitempty" yaml:"ignores,omitempty"`
		Rules struct {
			NoDefault bool     `json:"no_default,omitempty" yaml:"no_default,omitempty"`
			Add       []string `json:"add,omitempty" yaml:"add,omitempty"`
			Remove    []string `json:"remove,omitempty" yaml:"remove,omitempty"`
		} `json:"rules,omitempty" yaml:"rules,omitempty"`
		FileHeader struct {
			Path        string `json:"path,omitempty" yaml:"path,omitempty"`
			Content     string `json:"content,omitempty" yaml:"content,omitempty"`
			IsCommented bool   `json:"is_commented,omitempty" yaml:"is_commented,omitempty"`
		} `json:"file_header,omitempty" yaml:"file_header,omitempty"`
		JavaPackagePrefix string `json:"java_package_prefix,omitempty" yaml:"java_package_prefix,omitempty"`
		// devel-mode only
		AllowSuppression bool `json:"allow_suppression,omitempty" yaml:"allow_suppression,omitempty"`
	} `json:"lint,omitempty" yaml:"lint,omitempty"`
	LintCfg *lint.LintConfig
}

func NewConfig() (ret *Config, err error) {
	ret = &Config{}
	return
}

func (c *Config) Output() (err error) {
	var (
		file *os.File
	)
	file, err = os.Create("idl.yaml")
	if err != nil {
		err = errors.WithMessagef(err, "create idl.yaml failed")
		return
	}

	_, err = file.WriteString(tmpl)
	if err != nil {
		err = errors.WithMessagef(err, "write idl.yaml failed")
	}
	return
}

func (ret *Config) Load() (err error) {
	f, err := ioutil.ReadFile("idl.yaml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(f, ret)
	if err != nil {
		return err
	}

	//支持环境变量
	absPath := []string{}
	for _, itr := range ret.Includes {
		tmp := os.ExpandEnv(itr)
		absPath = append(absPath, tmp)
	}

	ret.Includes = absPath
	ret.ImportPath = os.ExpandEnv(ret.ImportPath)
	ret.Includes = append(ret.Includes, ret.ImportPath)
	for _, itr := range ret.Includes {
		log.Println("include path:", itr)
	}
	ignoreIDToFilePaths := make(map[string][]string)
	for _, ignore := range ret.Lint.Ignores {
		id := strings.ToUpper(ignore.ID)
		for _, protoFilePath := range ignore.Files {
			if !filepath.IsAbs(protoFilePath) {
				protoFilePath = filepath.Join(ret.ImportPath, protoFilePath)
			}
			protoFilePath = filepath.Clean(protoFilePath)
			if _, ok := ignoreIDToFilePaths[id]; !ok {
				ignoreIDToFilePaths[id] = make([]string, 0)
			}
			ignoreIDToFilePaths[id] = append(ignoreIDToFilePaths[id], protoFilePath)
		}
	}
	ret.LintCfg = &lint.LintConfig{
		IncludeIDs:          strs.SortUniqModify(ret.Lint.Rules.Add, strings.ToUpper),
		ExcludeIDs:          strs.SortUniqModify(ret.Lint.Rules.Remove, strings.ToUpper),
		Group:               strings.ToLower(ret.Lint.Group),
		NoDefault:           true,
		IgnoreIDToFilePaths: ignoreIDToFilePaths,
		//FileHeader:          fileHeader,
		JavaPackagePrefix: ret.Lint.JavaPackagePrefix,
		AllowSuppression:  false,
	}

	return nil
}
