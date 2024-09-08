package main

import (
	"testing"
)

func TestNameWithDiacritics(t *testing.T) {
	wanted := "Indio_forte_pummmmmmmmmm.mp3"
	input := "13. Índio forte pummmmmmmmmm - Fernando Rocha - vol.4.mp3"
	name := NormalizeClipName(input)

	if name != wanted {
		t.Fatalf("`NormalizeClipName(%q) = %q, wanted %q`", input, name, wanted)
	}
}

func TestNameWithCommas(t *testing.T) {
	wanted := "Barak_Obama,_Sarcosie,_Socrates.mp3"
	input := "11. Barak Obama, Sarcosie, Sócrates - Fernando Rocha - vol.7.mp3"
	name := NormalizeClipName(input)

	if name != wanted {
		t.Fatalf("`NormalizeClipName(%q) = %q, wanted %q`", input, name, wanted)
	}
}
