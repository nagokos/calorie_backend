package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type ConfigList struct {
	DbName    string `yaml:"name"`
	SQLDriver string `yaml:"driver"`
	Port      string `yaml:"port"`
}

var Config ConfigList

func init() {
	cfg, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Printf("err: %v", err)
	}

	var c map[string]ConfigList
	err = yaml.Unmarshal(cfg, &c)
	if err != nil {
		log.Printf("err: %v", err)
	}

	Config = ConfigList{
		DbName:    c["db"].DbName,
		SQLDriver: c["db"].SQLDriver,
		Port:      c["web"].Port,
	}
}
