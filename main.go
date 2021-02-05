package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/fhodun/anot/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.InitLogConfig()
	config := config.GetConfig()

	dg, err := discordgo.New("Bot " + config.DiscordToken)
	if err != nil {
		log.Fatal("Discord session creation failed")
	}
	err = dg.Open()
	if err != nil {
		log.Fatal("Unsuccessful opening connection")
	}

	defer dg.Close()

	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	log.Info("Bot is now running")
	select {}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || !strings.HasPrefix(m.Content, ">") {
		return
	}

	args := strings.Split(m.Content[(len(config.GetConfig().Prefix)):], ">")
	
	if len(args) < 3 {
		s.ChannelMessageDelete(m.ChannelID, m.ID)
		s.ChannelMessageSend(m.ChannelID, "Too few arguments")
		return
	}

	cmd := strings.Trim(args[0], " ")
	arg := args[1]
	msg := args[2]

	if cmd == "anon" {
		s.ChannelMessageDelete(m.ChannelID, m.ID)
		if arg != "" {
			s.ChannelMessageSend(arg, msg)
		} else {
			s.ChannelMessageSend(m.ChannelID, msg)
		}
	}
}
