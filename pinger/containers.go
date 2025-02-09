package main

import (
	"context"
	"log"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// ContainerStatus описывает сведения о контейнере для pinger.
type ContainerStatus struct {
	ID         string `json:"id"`
	IPAddress  string `json:"ip_address"`
	Status     string `json:"status"`
	Health     string `json:"health"`
	ProcessPID int    `json:"process_pid"`
}

// GetContainerStatuses получает список всех контейнеров через Docker API.
func GetContainerStatuses() ([]ContainerStatus, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Printf("[ERROR] Ошибка создания Docker клиента: %v", err)
		return nil, err
	}
	defer cli.Close()

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		log.Printf("[ERROR] Ошибка получения списка контейнеров: %v", err)
		return nil, err
	}

	var statuses []ContainerStatus
	for _, c := range containers {
		inspect, err := cli.ContainerInspect(context.Background(), c.ID)
		if err != nil {
			log.Printf("[WARNING] Ошибка инспекции контейнера %s: %v", c.ID, err)
			continue
		}

		ipAddress := "N/A"
		if inspect.NetworkSettings != nil && inspect.NetworkSettings.Networks != nil {
			for netName, network := range inspect.NetworkSettings.Networks {
				if network.IPAddress != "" {
					ipAddress = network.IPAddress
					log.Printf("[INFO] Контейнер %s -> Сеть: %s, IP: %s", c.ID[:12], netName, ipAddress)
					break
				}
			}
		} else {
			log.Printf("[WARNING] У контейнера %s нет сетевых интерфейсов", c.ID)
		}

		healthStatus := "unknown"
		if inspect.State.Health != nil {
			healthStatus = inspect.State.Health.Status
		}

		statuses = append(statuses, ContainerStatus{
			ID:         c.ID[:12],
			IPAddress:  ipAddress,
			Status:     inspect.State.Status,
			Health:     healthStatus,
			ProcessPID: inspect.State.Pid,
		})
	}

	log.Printf("[INFO] Всего обработано %d контейнеров", len(statuses))
	return statuses, nil
}
