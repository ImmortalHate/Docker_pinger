package main

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/streadway/amqp"
)

const queueName = "containers_queue"

var (
	rabbitConn    *amqp.Connection
	rabbitChannel *amqp.Channel
	once          sync.Once
)

func initRabbitMQ() {
	var err error
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		url = "amqp://guest:guest@broker:5672/"
	}
	rabbitConn, err = amqp.Dial(url)
	if err != nil {
		log.Fatalf("[ERROR] Ошибка подключения к RabbitMQ: %v", err)
	}

	rabbitChannel, err = rabbitConn.Channel()
	if err != nil {
		log.Fatalf("[ERROR] Ошибка создания канала RabbitMQ: %v", err)
	}

	_, err = rabbitChannel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("[ERROR] Ошибка объявления очереди RabbitMQ: %v", err)
	}
}

func SendMessage(result PingResult) {
	once.Do(initRabbitMQ)
	data, err := json.Marshal(result)
	if err != nil {
		log.Printf("[ERROR] Ошибка маршалинга JSON: %v", err)
		return
	}

	err = rabbitChannel.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         data,
			DeliveryMode: amqp.Persistent,
		},
	)
	if err != nil {
		log.Printf("[ERROR] Ошибка отправки сообщения: %v", err)
	} else {
		log.Printf("[INFO] Сообщение отправлено: %s", data)
	}
}
