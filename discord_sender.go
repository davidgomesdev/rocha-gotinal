package main

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

const username = "Fernando Rocha"
const avatarUrl = "https://i.ibb.co/QHHPPH4/roch1.jpg"

type WebhookPayload struct {
	Username string
	Avatar_url string
}

func SendClip(clip Clip, webhookUrl string) {
	file, _ := os.Open(clip.filePath)
  defer file.Close()

	body := &bytes.Buffer{}
  
	writer := multipart.NewWriter(body)
  part, _ := writer.CreateFormFile("files[0]", clip.name)
  io.Copy(part, file)
  writer.Close()

	payload, _ := json.Marshal(WebhookPayload{username, avatarUrl})

	writer.WriteField("payload_json", string(payload))

  r, _ := http.NewRequest("POST", webhookUrl, body)
  r.Header.Add("Content-Type", writer.FormDataContentType())
  client := &http.Client{}
  client.Do(r)
}
