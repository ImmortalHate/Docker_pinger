package main

import (
	"log"
	"time"
)

type PingResult struct {
	IPAddress   string    `json:"ip_address"`
	Status      string    `json:"status"`
	LastChecked time.Time `json:"last_checked"`
	PingTime    int64     `json:"ping_time"`
}

func main() {
	log.Println("Запуск pinger сервиса...")
	config := LoadConfig()

	if config.BackendURL == "" {
		config.BackendURL = "http://localhost:8080"
	}

	for {
		log.Println("Получаем список контейнеров...")
		containers, err := GetContainerStatuses()
		if err != nil {
			log.Printf("Ошибка получения контейнеров: %v", err)
			time.Sleep(config.PingInterval)
			continue
		}

		log.Printf("Обнаружено %d контейнеров", len(containers))

		for _, c := range containers {
			finalStatus := "OK"
			if c.Status != "running" || c.Health == "unhealthy" {
				finalStatus = "FAIL"
			}

			result := PingResult{
				IPAddress:   c.IPAddress,
				Status:      finalStatus,
				LastChecked: time.Now(),
				PingTime:    0,
			}

			log.Printf("Отправляем данные для контейнера %s: %+v", c.ID, result)
			if config.UseBroker {
				SendMessage(result)
			} else {
				sendDirect(result, config.BackendURL)
			}
		}

		log.Printf("Ждем %v перед следующим циклом...", config.PingInterval)
		time.Sleep(config.PingInterval)
	}
}
