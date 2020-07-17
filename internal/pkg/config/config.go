package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	yamlPath = "./config.yaml"
)

func New() *Config {
	return &Config{}
}

func (c *Config) ReadYaml() {
	b, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		fmt.Println(err)
	}

	if err = yaml.Unmarshal(b, c); err != nil {
		fmt.Println(err)
	}
}

func (c *Config) Log() {
	if b, err := yaml.Marshal(c); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(b))
	}
}

type Config struct {
	Port    string `yaml:"port"`
	GinMode string `yaml:"ginMode"`
	Mongo   struct {
		Database string `yaml:"db"`
		URI      string `yaml:"uri"`
	} `yaml:"mongo"`
}
