package models

import "time"

// ContainerStatus – модель для хранения состояния контейнера в БД.
type ContainerStatus struct {
	ID                 uint      `gorm:"primaryKey" json:"id"`
	IPAddress          string    `gorm:"uniqueIndex" json:"ip_address"`
	PingTime           int64     `json:"ping_time"`
	LastSuccessAttempt time.Time `json:"last_success_attempt"`
}
