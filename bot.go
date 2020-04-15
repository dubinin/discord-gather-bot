package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dubinin/discord-gather-bot/command"
	"github.com/dubinin/discord-gather-bot/gather"
	"github.com/bwmarrin/discordgo"

	log "github.com/Sirupsen/logrus"
)

// Token is Discord auth token
var Token = fmt.Sprintf("Bot %s", os.Getenv("BOT_TOKEN"))

var Gather = gather.Gather{
	Games:    make(map[string]gather.Gatherable),
	ChanelID: os.Getenv("BOT_CHANNEL"),
	Admins:   strings.Split(os.Getenv("BOT_ADMINS"), ","),
}

func main() {
	dg, err := discordgo.New(Token)
	if err != nil {
		log.Panicf("Error creating Discord session, %v", err)
		return
	}

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		log.Panicf("Error opening connection, %v", err)
		return
	}

	log.Info("Bot is now running.  Press CTRL-C to exit.")
	log.Infof("Use token %s", Token)
	log.Infof("Admins are: %s", Gather.Admins)
	// Simple way to keep program running until CTRL-C is pressed.
	<-make(chan struct{})
	return
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(Gather.ChanelID) == 0 || m.ChannelID == Gather.ChanelID {
		if strings.HasPrefix(m.Content, "!") {
			command.New(s, m, &Gather).Execute()
		}
	}
}
