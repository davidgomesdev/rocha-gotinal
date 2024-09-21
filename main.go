package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	channelId := os.Getenv("CHANNEL_ID")

	if token == "" || channelId == "" {
		log.Fatal("You need to provide BOT_TOKEN and CHANNEL_ID environment variables.")
	}

	webhookFile, err := os.ReadFile(".webhook_url")
	exitOnErr(err)

	webhookUrl := string(webhookFile)

	if webhookUrl == "" {
		log.Fatal("You need to have a .webhook_url file.")
	}

	discord, err := discordgo.New(fmt.Sprint("Bot ", token))
	exitOnErr(err)

	sentClips := getSentClips(discord, channelId)

	log.Println("Excluding", len(sentClips), "messages")

	clip := GetRandomClip(sentClips)

	log.Println("Sending", clip.name)

	SendClip(clip, webhookUrl)
}

func exitOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
