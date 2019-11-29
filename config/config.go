package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
)

// Cfg - variable for storing loaded configuration
var Cfg Config

// Config - configuration structure
type Config struct {
	Port     string   `json:"port", envconfig:"PORT"`
	Host     string   `json:"host", envconfig:"HOST"`
	Database Database `json:"database"`
}

// Database - database configuration structure
type Database struct {
	DBHost   string `json:"host", envconfig:"DB_HOST"`
	DBPort   string `json:"port", envconfig:"DB_PORT"`
	Username string `json:"user", envconfig:"DB_USERNAME"`
	Password string `json:"pass", envconfig:"DB_PASSWORD"`
	DBName   string `json:"dbname", envconfig:"DB_NAME"`
}

func readFile(cfg *Config) {
	f, err := os.Open("config/config.json")
	if err != nil {
		fmt.Println(err)
	}

	decoder := json.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Panic(err)
	}
}

func readEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		fmt.Println(err)
	}
}

// SetConfig - function to load configuration
func SetConfig() {
	readFile(&Cfg)
	readEnv(&Cfg)
}
