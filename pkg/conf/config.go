package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
	Includes []string `yaml:"includes"`
	Excludes []string `yaml:"excludes"`
	Depends  []string `yaml:"depends"`
	Generate Generate `yaml:"generate"`
}

var IDLConfig = Config{}

func Init() error {
	f, err := ioutil.ReadFile("idl.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(f, &IDLConfig)
	if err != nil {
		return err
	}
	return nil
}
