package main

import (
	"github.com/bwmarrin/discordgo"
)

const fetchSize = 100

func getSentClips(discord *discordgo.Session, channelId string) []string {
	var totalClipsSent []string

	msgs, err := discord.ChannelMessages(channelId, fetchSize, "", "", "")
	exitErr(err)

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
