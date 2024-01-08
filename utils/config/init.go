package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var C = new(config)

func init() {
	// Load configurations to struct
	yml, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("UNABLE TO READ YAML CONFIGURATION FILE")
	}
	if err := yaml.Unmarshal(yml, C); err != nil {
		log.Fatal("UNABLE TO PARSE YAML CONFIGURATION FILE")
	}
}
