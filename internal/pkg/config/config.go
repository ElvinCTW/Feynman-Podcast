package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
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
	fmt.Println("port", c.Port)
	fmt.Println("ginMode", c.GinMode)
	fmt.Println("mongo database", c.Mongo.Database)
	fmt.Println("mongo uri", len(c.Mongo.URI))
	fmt.Println("awsKey", len(c.Upload.AwsAccessKey))
	fmt.Println("awsSecret", len(c.Upload.AwsAccessSecret))
	fmt.Println("awsRegion", c.Upload.AwsRegion)
	fmt.Println("awsBucket", c.Upload.AwsBucket)
}

func (c *Config) ReadEnv() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("not used .env")
	}

	c.Mongo.URI = os.Getenv(c.Mongo.URI)
	c.Upload.AwsAccessKey = os.Getenv(c.Upload.AwsAccessKey)
	c.Upload.AwsAccessSecret = os.Getenv(c.Upload.AwsAccessSecret)
}

type Config struct {
	Port    string `yaml:"port"`
	GinMode string `yaml:"ginMode"`
	Mongo   struct {
		Database string `yaml:"db"`
		URI      string `yaml:"uri"`
	} `yaml:"mongo"`
	Upload struct {
		AwsAccessKey    string `yaml:"awsAccessKey"`
		AwsAccessSecret string `yaml:"awsAccessSecret"`
		AwsRegion       string `yaml:"awsRegion"`
		AwsBucket       string `yaml:"awsBucket"`
	}
}
