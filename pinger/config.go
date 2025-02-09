package main

import (
	"os"
	"time"
)

type Config struct {
	BackendURL   string
	RabbitMQURL  string
	PingInterval time.Duration
	UseBroker    bool
}

func LoadConfig() Config {
	useBroker := false
	if os.Getenv("USE_BROKER") == "true" {
		useBroker = true
	}
	return Config{
		BackendURL:   os.Getenv("BACKEND_URL"),
		RabbitMQURL:  os.Getenv("RABBITMQ_URL"),
		PingInterval: 10 * time.Second,
		UseBroker:    useBroker,
	}
}
