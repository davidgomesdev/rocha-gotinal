package main

import (
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"

	"github.com/thoas/go-funk"
	"golang.org/x/text/unicode/norm"
)

const clipsFolder = "clips/"

const extract_clip_name_regex = "\\d\\d\\. (.+?) -.+\\.mp3"
const sanitize_regex = "[^a-zA-Z0-9 -.]+"

type Clip struct {
	name     string
	filePath string
}

func GetRandomClip(sentClips []string) Clip {
	clips := getUnsentClips(sentClips)

	log.Println("There are", len(clips), "unsent clips")

	return clips[rand.Intn(len(clips))]
}

func getUnsentClips(sentClips []string) []Clip {
	clipFiles, err := os.ReadDir(clipsFolder)
	exitErr(err)

	var clips []Clip

	for _, clip := range clipFiles {
		clipName := clip.Name()
		isIntroClip := strings.HasPrefix(clipName, "01")

		if isIntroClip {
			continue
		}

		if normalizedClipName := normalizeClipName(clipName); normalizedClipName != "" {
			if !funk.Contains(sentClips, func(sentClipName string) bool {
				return sentClipName == normalizedClipName
			}) {
				var sb strings.Builder

				sb.WriteString(clipsFolder)
				sb.WriteString(clipName)

				clips = append(clips, Clip{normalizedClipName, sb.String()})
			}
		}
	}

	return clips
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

	var sb strings.Builder

	sb.WriteString(strings.ReplaceAll(strings.TrimSpace(text), " ", "_"))
	sb.WriteString(".mp3")

	return sb.String()
}