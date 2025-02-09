package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func sendDirect(result PingResult, backendURL string) {
	data, err := json.Marshal(result)
	if err != nil {
		log.Printf("Ошибка сериализации JSON: %v", err)
		return
	}

	url := backendURL + "/api/status"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Printf("Ошибка создания HTTP запроса: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Ошибка отправки HTTP запроса: %v", err)
		return
	}
	defer resp.Body.Close()
	log.Printf("Прямая отправка завершена, статус: %s", resp.Status)
}
