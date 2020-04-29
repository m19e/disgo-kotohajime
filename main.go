package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	Token string
	BotName string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Token = os.Getenv("TOKEN")
	BotName = os.Getenv("CLIENT_ID")
}

func main() {
	// Create a new Discord session.
	dg, err := discordgo.New("Bot " + Token); if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// Register the messageCreate func as a callback
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open(); if err != nil {
		fmt.Println("Error opening connection: ", err)
		return
	}

	// Wait here until Ctrl-C or other term signal is received.
	fmt.Println("Bot is now running. Press Ctrl-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}


// This func will be called every time a new message
// is created on any channel that bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Get channel where the message created.
	c, err := s.State.Channel(m.ChannelID); if err != nil {
		log.Println("Error getting channel: ", err)
	}

	// Ignore all messages created by the bot itself.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// If the message is "ping" reply with "Pong!".
	if m.Content == "ping" {
		sendMsg(s, c, "Pong!")
	}

	// If the message is "pong" reply with "Ping!".
	if m.Content == "pong" {
		sendMsg(s, c, "Ping!")
	}
}

func sendMsg(s *discordgo.Session, c *discordgo.Channel, msg string) {
	log.Println(">>> " + msg)
	_, err := s.ChannelMessageSend(c.ID, msg); if err != nil {
		log.Println("Error sending message: ", err)
	}
}
