package database

import (
	"fmt"
	"log"
	"os"
	"vk-pinger/backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectDB открывает соединение с БД, выполняет миграцию и возвращает объект GORM.
func ConnectDB() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	// Используем стандартный логгер GORM с подробностями
	db, err := openDB(dsn)
	if err != nil {
		log.Fatalf("[ERROR] Ошибка подключения к БД: %v", err)
	}

	if err := db.AutoMigrate(&models.ContainerStatus{}); err != nil {
		log.Fatalf("[ERROR] Ошибка миграции: %v", err)
	}
	log.Println("[INFO] Миграция БД завершена успешно.")
	return db
}

// openDB оборачивает gorm.Open с предварительными настройками.
func openDB(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
