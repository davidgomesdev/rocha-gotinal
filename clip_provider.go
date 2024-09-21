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
			normalizedClipName := NormalizeClipName(path)

			if isClipUnsent(normalizedClipName, sentClips) {
				unsentClips = append(unsentClips, Clip{normalizedClipName, path})
			}
		}

		return nil
	})
	exitOnErr(err)

	return unsentClips
}

func isClipUnsent(normalizedClipName string, sentClips []string) bool {
	isIntroClip := strings.HasPrefix(normalizedClipName, "01")

	if isIntroClip {
		return false
	}

	if !funk.Contains(sentClips, func(sentClipName string) bool {
		return sentClipName == normalizedClipName
	}) {
		return true
	}

	return false
}
