package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var config *AppConfig

func Config() *AppConfig {

	if nil != config {
		return config
	}

	config = &AppConfig{}

	yamlFile, err := ioutil.ReadFile("./app-config.yaml")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return config
}

type AppConfig struct {
	Host  string `yaml:"host"`
	DbUrl string `yaml:"db-url"`
	Log Log `yaml:"log"`
}

type Log struct {
	Path   string `yaml:"path"`
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}
