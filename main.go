package main

import (
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/fhodun/anot/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	config := config.GetConfig()

	log.SetFormatter(&log.TextFormatter{
		ForceColors: true,
	})
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

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

	mContent := m.Content[(len(config.GetConfig().Prefix)):]
	args := strings.Split(mContent, ">")
	cmd := strings.Trim(args[0], " ")
	arg := args[1]

	if cmd == "anon" {
		s.ChannelMessageDelete(m.ChannelID, m.ID)
		s.ChannelMessageSend(m.ChannelID, arg)
	}
}
