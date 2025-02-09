package api

import (
	"net/http"
	"time"

	"vk-pinger/backend/repository"

	"github.com/gin-gonic/gin"
	gojson "github.com/goccy/go-json"
)

// RegisterRoutes регистрирует HTTP API-обработчики.
func RegisterRoutes(router *gin.Engine, repo repository.PostgresRepositoryInterface) {
	router.GET("/api/status", func(c *gin.Context) {
		statuses, err := repo.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		data, err := gojson.Marshal(statuses)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Data(http.StatusOK, "application/json", data)
	})

	router.POST("/api/status", func(c *gin.Context) {
		var status repository.ContainerStatus
		if err := c.ShouldBindJSON(&status); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if status.LastSuccessAttempt.IsZero() {
			status.LastSuccessAttempt = time.Now()
		}
		if err := repo.Save(status); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusCreated)
	})
}
