package conf

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type Plugin struct {
	Name   string `yaml:"name"`
	Flags  string `yaml:"flags"`
	Type   string `yaml:"type"`
	Output string `yaml:"output"`
}

type GoOptions struct {
	ExtraModifiers map[string]string `yaml:"extra_modifiers"`
	Modifier       string            `yaml:"modifier"`
}

type Generate struct {
	GoOptions GoOptions `yaml:"go_options"`
	Plugins   []Plugin  `yaml:"plugins"`
	Output    string    `yaml:"output"`
}

type Config struct {
	ImportPath string              `yaml:"import_path"`
	GoModule   string              `yaml:"go_module"`
	Protos     []string            `yaml:"protos"`
	Ignore     []string            `yaml:"ignore"`
	IngoreMap  map[string]struct{} `yaml:"-"`
	Includes   []string            `yaml:"includes"`
	Generate   Generate            `yaml:"generate"`
	Lint       struct {
		Rules struct {
			Enable  []string `json:"enable,omitempty" yaml:"enable,omitempty"`
			Disable []string `json:"disable,omitempty" yaml:"disable,omitempty"`
		} `json:"rules,omitempty" yaml:"rules,omitempty"`
	} `json:"lint,omitempty" yaml:"lint,omitempty"`
}

func NewConfig() (ret *Config) {
	ret = &Config{
		IngoreMap: make(map[string]struct{}),
	}
	return
}

var configFileName = "bup.yaml"

func (c *Config) Output() (err error) {
	var (
		file *os.File
	)
	file, err = os.Create(configFileName)
	if err != nil {
		err = errors.WithMessagef(err, "create bup.yaml failed")
		return
	}

	_, err = file.WriteString(tmpl)
	if err != nil {
		err = errors.WithMessagef(err, "write bup.yaml failed")
	}
	return
}

func (c *Config) Load() (err error) {
	f, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(f, c)
	if err != nil {
		return err
	}

	// include path
	absPath := []string{}
	for _, itr := range c.Includes {
		tmp := os.ExpandEnv(itr)
		tmp = filepath.ToSlash(tmp)
		absPath = append(absPath, tmp)
	}
	c.Includes = absPath

	//ignore file
	for _, itr := range c.Ignore {
		c.IngoreMap[itr] = struct{}{}
	}

	//import path
	c.ImportPath = filepath.FromSlash(os.ExpandEnv(c.ImportPath))
	c.Includes = append(c.Includes, c.ImportPath)
	for _, itr := range c.Includes {
		log.Println("include path:", itr)
	}

	//output path
	for _, itr := range c.Generate.Plugins {
		itr.Output = os.ExpandEnv(itr.Output)
	}
	c.Generate.Output = os.ExpandEnv(c.Generate.Output)

	return nil
}
