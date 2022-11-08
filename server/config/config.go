package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type config struct {
	Server server `yaml:"server"`
}

type server struct {
	Port     int    `yaml:"port"`
	BasePath string `yaml:"basePath"`
}

func GetConfig() *config {
	configFile, err := ioutil.ReadFile("config/config.yaml")

	if err != nil {
		log.Fatalf("Fatal error reading in config.yaml: %v", err)
	}

	conf := &config{}

	err = yaml.Unmarshal(configFile, conf)

	if err != nil {
		log.Fatalf("Fatal error decoding config.yaml: %v", err)
	}

	return conf
}
