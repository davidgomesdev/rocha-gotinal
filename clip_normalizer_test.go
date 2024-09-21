package main

import (
	"testing"
)

func TestNameWithDiacritics(t *testing.T) {
	wanted := "Indio_forte_pummmmmmmmmm.mp3"
	input := "13. Índio forte pummmmmmmmmm - Fernando Rocha - vol.4.mp3"
	name := NormalizeClipName(input)

	if name != wanted {
		t.Fatalf("`NormalizeClipName(%q) = %q, WANTED %q`", input, name, wanted)
	}
}

func TestNameWithCommas(t *testing.T) {
	wanted := "Barak_Obama_Sarcosie_Socrates.mp3"
	input := "11. Barak Obama, Sarcosie, Sócrates - Fernando Rocha - vol.7.mp3"
	name := NormalizeClipName(input)

	if name != wanted {
		t.Fatalf("`NormalizeClipName(%q) = %q, WANTED %q`", input, name, wanted)
	}
}

func TestNameWithParenthesis(t *testing.T) {
	wanted := "Fazer_rir_e_um_bem_necessario_ate_ao_CD_8.mp3"
	input := "16. Fazer rir é um bem necessário (até ao CD 8) - Fernando Rocha - vol.7.mp3"
	name := NormalizeClipName(input)

	if name != wanted {
		t.Fatalf("`NormalizeClipName(%q) = %q, WANTED %q`", input, name, wanted)
	}
}

func TestNameWithFullPath(t *testing.T) {
	wanted := "Indio_forte_pummmmmmmmmm.mp3"
	input := "/clips/Volume 4/13. Índio forte pummmmmmmmmm - Fernando Rocha - vol.4.mp3"
	name := NormalizeClipName(input)

	if name != wanted {
		t.Fatalf("`NormalizeClipName(%q) = %q, WANTED %q`", input, name, wanted)
	}
}
