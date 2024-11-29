package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

const username = "Fernando Rocha"
const avatarUrl = "https://i.ibb.co/QHHPPH4/roch1.jpg"

type WebhookFormPayload struct {
	Username  string `json:"username"`
	AvatarUrl string `json:"avatar_url"`
}

type WebhookJsonPayload struct {
	Username  string `json:"username"`
	AvatarUrl string `json:"avatar_url"`
	Content   string `json:"content"`
}

func SendClip(clip Clip, webhookUrl string) {
	file, _ := os.Open(clip.filePath)
	defer file.Close()

	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("files[0]", clip.name)
	io.Copy(part, file)

	payload, _ := json.Marshal(WebhookFormPayload{username, avatarUrl})

	writer.WriteField("payload_json", string(payload))
	writer.Close()

	r, _ := http.NewRequest("POST", webhookUrl, body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	client.Do(r)
}

func SendMessage(message string, webhookUrl string) {
	payload, _ := json.Marshal(WebhookJsonPayload{username, avatarUrl, message})

	r, _ := http.Post(webhookUrl, "application/json", bytes.NewBuffer(payload))
	responseBytes, _ := io.ReadAll(r.Body)

	log.Println("Got", r.Status, string(responseBytes), "for sending the message:", message)
}
