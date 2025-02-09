package repository

import "time"

// ContainerStatus описывает данные состояния контейнера для бизнес-логики.
type ContainerStatus struct {
	ID                 uint      `json:"id"`
	IPAddress          string    `json:"ip_address"`
	PingTime           int64     `json:"ping_time"`
	LastSuccessAttempt time.Time `json:"last_success_attempt"`
}
