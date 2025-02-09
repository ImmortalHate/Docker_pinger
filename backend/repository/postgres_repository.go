package repository

import (
	"vk-pinger/backend/models"

	"gorm.io/gorm"
)

type PostgresRepositoryInterface interface {
	GetAll() ([]ContainerStatus, error)
	Save(status ContainerStatus) error
}

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) PostgresRepositoryInterface {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) GetAll() ([]ContainerStatus, error) {
	var modelStatuses []models.ContainerStatus
	if err := r.db.Find(&modelStatuses).Error; err != nil {
		return nil, err
	}
	var statuses []ContainerStatus
	for _, s := range modelStatuses {
		statuses = append(statuses, ContainerStatus{
			ID:                 s.ID,
			IPAddress:          s.IPAddress,
			PingTime:           s.PingTime,
			LastSuccessAttempt: s.LastSuccessAttempt,
		})
	}
	return statuses, nil
}

func (r *PostgresRepository) Save(status ContainerStatus) error {
	var existing models.ContainerStatus
	err := r.db.Where("ip_address = ?", status.IPAddress).First(&existing).Error
	if err != nil {
		newStatus := models.ContainerStatus{
			IPAddress:          status.IPAddress,
			PingTime:           status.PingTime,
			LastSuccessAttempt: status.LastSuccessAttempt,
		}
		return r.db.Create(&newStatus).Error
	}
	existing.PingTime = status.PingTime
	existing.LastSuccessAttempt = status.LastSuccessAttempt
	return r.db.Save(&existing).Error
}
