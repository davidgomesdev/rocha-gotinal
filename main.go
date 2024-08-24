package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/bwmarrin/discordgo"
)

const fetchSize = 100

func main() {
	token := os.Getenv("BOT_TOKEN")
	channelId := os.Getenv("CHANNEL_ID")
	discord, err := discordgo.New(fmt.Sprint("Bot ", token))
	exitErr(err)

	msgs, err := discord.ChannelMessages(channelId, fetchSize, "", "", "")
	exitErr(err)

	totalMsgs := make([]discordgo.Message, len(msgs))

	for ok := true; ok; ok = len(msgs) > 0 {
		for i := range msgs {
			msg := msgs[i]

			if len(msg.Attachments) > 0 {
				fmt.Println(msg.Attachments[0].Filename)
			}
		}

		for item := range msgs {
			if msgs[item] != nil {
				totalMsgs = append(totalMsgs, *msgs[item])
			}
		}

		newMsgs, err := discord.ChannelMessages(channelId, fetchSize, msgs[len(msgs)-1].ID, "", "")
		exitErr(err)
		msgs = newMsgs
	}

	for _, msg := range totalMsgs {
		if len(msg.Attachments) > 0 {
			fmt.Println(msg.Attachments[0].Filename)
		}
	}

	fmt.Println("---------------------------")
	fmt.Println(fmt.Sprint("Got ", len(totalMsgs), " messages"))
}

func exitErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		debug.PrintStack()
		os.Exit(1)
	}
}
