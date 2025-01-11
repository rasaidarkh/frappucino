package service

import (
	"frappuccino/internal/models"
	"log/slog"
)

type InventoryRepository interface {
	GetAll() ([]models.Inventory, error)
}

type Service struct {
	Repo   InventoryRepository
	logger *slog.Logger
}

type InventoryService interface {
	GetAll() ([]models.Inventory, error)
}

func NewInventoryService(repo InventoryRepository, logger *slog.Logger) *Service {
	return &Service{Repo: repo, logger: logger}
}

func (s *Service) GetAll() ([]models.Inventory, error) {
	return s.Repo.GetAll()
}
