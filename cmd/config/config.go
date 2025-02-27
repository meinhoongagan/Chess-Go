package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		URL string `yaml:"url"`
	} `yaml:"database"`
}

var AppConfig Config

func LoadConfig() {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("❌ Error reading config file: %v", err)
	}

	err = yaml.Unmarshal(file, &AppConfig)
	if err != nil {
		log.Fatalf("❌ Error parsing YAML: %v", err)
	}

	log.Println("✅ Config loaded successfully")
}
