package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var config *AppConfig

func GetConfig() *AppConfig {

	if nil != config {
		return config
	}

	yamlFile, err := ioutil.ReadFile("./app-config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		panic(err)
	}

	return config
}

type AppConfig struct {
	Host  string `yaml:"host"`
	DbUrl string `yaml:"db-url"`
}
