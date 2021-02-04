package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/fhodun/anot/config"
)

func main() {
	config := config.GetConfig()
	dg, err := discordgo.New("Bot " + config.DiscordToken)
	if err != nil {
		log.Fatal("[FATAL ERROR] Discord session creation failed, ", err)
	}
	err = dg.Open()
	if err != nil {
		log.Fatal("[FATAL ERROR] Unsuccessful opening connection, ", err)
	}

	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	println("[INFO] Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if (m.Author.ID == s.State.User.ID || !strings.HasPrefix(m.Content, ">")) { return }

	mContent := m.Content[(len(config.GetConfig().Prefix)):]
	var args []string
	args = strings.Split(mContent, ">")
	
	cmd := strings.Trim(args[0], " ")
	arg := args[1]

	if cmd == "anon" {
		s.ChannelMessageDelete(m.ChannelID, m.ID)
		s.ChannelMessageSend(m.ChannelID, arg)
	}
}