package main

import (
	"testing"
)

func TestClipPresentInSentClipsIsNotUnsent(t *testing.T) {
	input := "/rocha-gotinal/clips/Volume 7/16. Fazer rir é um bem necessário (até ao CD 8) - Fernando Rocha - vol.7.mp3"
	inputSent := NormalizeClipName("16. Fazer rir é um bem necessário (até ao CD 8) - Fernando Rocha - vol.7.mp3")
	dummySent := NormalizeClipName("05. Sem braços, sem pernas - Fernando Rocha - vol.5.mp3")
	sentClips := []string{inputSent, dummySent}

	isUnsent := isClipUnsent(NormalizeClipName(input), sentClips)

	if isUnsent {
		t.Fatalf("clip should not be unsent")
	}
}

func TestClipNotPresentInSentClipsIsUnsent(t *testing.T) {
	input := "/rocha-gotinal/clips/Volume 7/16. Fazer rir é um bem necessário (até ao CD 8) - Fernando Rocha - vol.7.mp3"
	dummySent := NormalizeClipName("05. Sem braços, sem pernas - Fernando Rocha - vol.5.mp3")
	sentClips := []string{dummySent}

	isUnsent := isClipUnsent(input, sentClips)

	if !isUnsent {
		t.Fatalf("clip should be unsent")
	}
}
