package conf

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Plugin struct {
	Name   string `yaml:"name"`
	Output string `yaml:"output"`
	Flags  string `yaml:"flags"`
	Type   string `yaml:"type"`
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
		Rules struct {
			Enable  []string `json:"enable,omitempty" yaml:"enable,omitempty"`
			Disable []string `json:"disable,omitempty" yaml:"disable,omitempty"`
		} `json:"rules,omitempty" yaml:"rules,omitempty"`
	} `json:"lint,omitempty" yaml:"lint,omitempty"`
}

func NewConfig() (ret *Config) {
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
		tmp = filepath.ToSlash(tmp)
		absPath = append(absPath, tmp)
	}

	ret.Includes = absPath
	ret.ImportPath = filepath.FromSlash(os.ExpandEnv(ret.ImportPath))
	ret.Includes = append(ret.Includes, ret.ImportPath)
	for _, itr := range ret.Includes {
		log.Println("include path:", itr)
	}
	return nil
}
