package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
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
}

func NewConfig() (ret *Config, err error) {
	ret = &Config{}
	f, err := ioutil.ReadFile("idl.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(f, &ret)
	if err != nil {
		return nil, err
	}

	//支持环境变量
	absPath := []string{}
	for _, itr := range ret.Includes {
		tmp := os.ExpandEnv(itr)
		log.Println("include path:", tmp)
		absPath = append(absPath, tmp)
	}

	ret.Includes = absPath
	ret.ImportPath = os.ExpandEnv(ret.ImportPath)

	return
}
