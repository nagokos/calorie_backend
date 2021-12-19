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
	Log       string `yaml:"log"`
	DBConf    string `yaml:"dbconf"`
}

var Config ConfigList

func init() {
	cfg, err := ioutil.ReadFile("config/config.yml")
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
		Log:       c["web"].Log,
		DBConf:    c["db"].DBConf,
	}
}
