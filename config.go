package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type generate struct {
	OutPath string `yaml:"outPath"`
	GoOptions struct{
		ExtraModifiers map[string]string `yaml:"extraModifiers"`
	}`yaml:"goOptions"`
}

type idlConfig struct {
	IDLPath string `yaml:"idlPath"`
	IncludePath []string `yaml:"includePath"`
	Excludes []string `yaml:"excludes"`
	ExcludesMap map[string]struct{}
	Depends []string `yaml:"depends"`
	Generate    generate `yaml:"generate"`
}

var IDLConfig = idlConfig{}

func Init() error {
	f, err := ioutil.ReadFile("idl.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(f, &IDLConfig)
	if err != nil {
		return err
	}
	for _, itr := range IDLConfig.Excludes {
		IDLConfig.ExcludesMap[itr] = struct{}{}
	}
	return nil
}
