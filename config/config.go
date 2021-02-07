package config

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"os"
)

// Config dupa
type Config struct {
	Discord struct {
		Token  string `json:"token"`
		Prefix string `json:"prefix"`
	} `json:"discord"`
}

// InitLogConfig dupa
func InitLogConfig() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(os.Stdout)
}

// LoadConfig dupa
func LoadConfig() *Config {
	config := &Config{}
	configFile, err := os.Open("config.json")
	defer configFile.Close()
	if err != nil {
		log.Fatal("Unsuccessful opening config file, ", err)
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
