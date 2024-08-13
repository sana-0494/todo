package cmd

import (
	"fmt"
	"log"
	"todo/configs"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

var k = koanf.New(".")

func getConfig(f string) configs.Config {

	if err := k.Load(file.Provider(f), yaml.Parser()); err != nil {
		log.Fatalf("error loading the config file: %v", err)
	}
	var cfg configs.Config
	err := k.UnmarshalWithConf("", &cfg, koanf.UnmarshalConf{FlatPaths: false})
	if err != nil {
		fmt.Println("error umarshalling the configs", err)
	}
	return cfg
}
