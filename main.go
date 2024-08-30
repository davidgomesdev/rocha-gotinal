package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/text/unicode/norm"
)

const fetchSize = 100

const extract_clip_name_regex = "\\d\\d\\. (.+?) -.+\\.mp3"
const sanitize_regex = "[^a-zA-Z0-9 -.]+"

func main() {
	token := os.Getenv("BOT_TOKEN")
	channelId := os.Getenv("CHANNEL_ID")

	if token == "" || channelId == "" {
		log.Fatal("You need to provide BOT_TOKEN and CHANNEL_ID environment variables.")
	}

	discord, err := discordgo.New(fmt.Sprint("Bot ", token))
	exitErr(err)

	msgs, err := discord.ChannelMessages(channelId, fetchSize, "", "", "")
	exitErr(err)

	sentClips := getSentClips(msgs, discord, channelId)

	clips, err := os.ReadDir("clips/")
	if err != nil {
		log.Fatal(err)
	}

	for _, clip := range clips {
		log.Println("normalized", normalizeClipName(clip.Name()))
	}

	log.Println("---------------------------")
	log.Println("Got", len(sentClips), "messages")
}

func getSentClips(msgs []*discordgo.Message, discord *discordgo.Session, channelId string) []string {
	var totalClipsSent []string

	for ok := true; ok; ok = len(msgs) > 0 {
		for _, msg := range msgs {
			if msg != nil && len(msg.Attachments) > 0 {
				if fileName := msg.Attachments[0].Filename; fileName != "" {
					totalClipsSent = append(totalClipsSent, msg.Attachments[0].Filename)
				}
			}
		}

		newMsgs, err := discord.ChannelMessages(channelId, fetchSize, msgs[len(msgs)-1].ID, "", "")
		exitErr(err)
		msgs = newMsgs
	}

	return totalClipsSent
}

func normalizeClipName(originalText string) string {
	var text = norm.NFC.String(originalText)

	var re = regexp.MustCompile(extract_clip_name_regex)

	match := re.FindStringSubmatch(originalText)

	if len(match) == 1 {
		return ""
	}

	text = match[1]
	re = regexp.MustCompile(sanitize_regex)
	text = re.ReplaceAllString(text, "")
	
	return strings.ReplaceAll(strings.TrimSpace(text), " ", "_")
}

func exitErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
