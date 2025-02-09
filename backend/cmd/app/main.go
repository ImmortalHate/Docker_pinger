package main

import (
	"fmt"
	"log"
	"os"
	"time"

	api "vk-pinger/backend/api/handlers"
	"vk-pinger/backend/database"
	"vk-pinger/backend/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func initDB() *repository.PostgresRepositoryInterface {
	db := database.ConnectDB()
	repo := repository.NewPostgresRepository(db)
	return &repo
}

func main() {
	log.Println("[INFO] Запуск backend сервиса...")
	db := database.ConnectDB()
	repo := repository.NewPostgresRepository(db)

	// Настройка Gin и CORS
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api.RegisterRoutes(router, repo)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)
	log.Printf("[INFO] Backend запущен на %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("[ERROR] Ошибка запуска сервера: %v", err)
	}
}
