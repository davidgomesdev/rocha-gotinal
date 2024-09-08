package main

import (
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/thoas/go-funk"
)

const clipsFolder = "clips/"

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
	var unsentClips []Clip

	err := filepath.WalkDir(clipsFolder, func(path string, _ os.DirEntry, err error) error {
		if err == nil && strings.HasSuffix(path, ".mp3") {
			clipName := path[strings.LastIndex(path, "/")+1:]

			isIntroClip := strings.HasPrefix(clipName, "01")

			if isIntroClip {
				return nil
			}

			if normalizedClipName := NormalizeClipName(clipName); normalizedClipName != "" {
				if !funk.Contains(sentClips, func(sentClipName string) bool {
					return sentClipName == normalizedClipName
				}) {
					var sb strings.Builder

					sb.WriteString(clipsFolder)
					sb.WriteString(clipName)

					unsentClips = append(unsentClips, Clip{normalizedClipName, sb.String()})
				}
			}
		}
		return nil
	})
	exitErr(err)

	return unsentClips
}
