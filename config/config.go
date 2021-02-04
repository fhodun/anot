package config

import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

// Config dupa
type Config struct {
    DiscordToken string
    Prefix       string
}

// GetConfig dupa
func GetConfig() *Config {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("[ERROR] Unsuccessful loading .env, ", err)
    }

    discordToken, discordTokenExists := os.LookupEnv("DISCORD_TOKEN")
    prefix, prefixExists := os.LookupEnv("PREFIX")
    
    if !discordTokenExists {
        log.Fatal("[FATAL ERROR] No discord token detected")
    }
    if !prefixExists {
        // TODO: need repair and make this better
        //println("[INFO] No prefix detected, default '>' will be used")
        prefix = ">"
    }

    config := &Config{
        DiscordToken: discordToken,
        Prefix:       prefix,
    }
    return config
}