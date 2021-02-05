package config

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

// Config dupa
type Config struct {
	DiscordToken string
	Prefix       string
}

// InitLogConfig dupa
func InitLogConfig() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:               true,
		DisableColors:             false,
		ForceQuote:                false,
		DisableQuote:              false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             false,
		TimestampFormat:           "",
		DisableSorting:            false,
		SortingFunc: func([]string) {
		},
		DisableLevelTruncation: false,
		PadLevelText:           false,
		QuoteEmptyFields:       false,
	})
	log.SetOutput(os.Stdout)
}

// GetConfig dupa
func GetConfig() *Config {
	InitLogConfig()
	err := godotenv.Load()
	if err != nil {
		log.Warn("Unsuccessful loading .env, ", err)
	}

	discordToken, discordTokenExists := os.LookupEnv("DISCORD_TOKEN")
	prefix, prefixExists := os.LookupEnv("PREFIX")

	if !discordTokenExists {
		log.Fatal("No discord token detected")
	}
	if !prefixExists {
		// TODO: need repair and make this better
		//log.Info("No prefix detected, default '>' will be used")
		prefix = ">"
	}

	config := &Config{
		DiscordToken: discordToken,
		Prefix:       prefix,
	}
	return config
}
