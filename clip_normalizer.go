package main

import (
	"regexp"
	"strings"

	"golang.org/x/text/unicode/norm"
)

const extract_clip_name_regex = "\\d\\d\\. (.+?) -.+\\.mp3"
const sanitize_regex = "[^a-zA-Z0-9 -.]+"

func NormalizeClipName(originalText string) string {
	var text = norm.NFD.String(originalText)
	var re = regexp.MustCompile(extract_clip_name_regex)

	match := re.FindStringSubmatch(text)

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
