package service

import (
	"frappuccino/internal/models"
	"log/slog"
)

type InventoryRepository interface {
	GetAll() ([]models.Inventory, error)
	GetElementById(InventoryId int) (models.Inventory, error)
	Delete(InventoryId int) error
	Put(item models.Inventory) error
	Post(item models.Inventory) error
}

type InventoryService struct {
	Repo   InventoryRepository
	logger *slog.Logger
}

func (s *InventoryService) Put(item models.Inventory) error {
	//TODO implement me
	panic("implement me")
}

func (s *InventoryService) Post(item models.Inventory) error {
	//TODO implement me
	panic("implement me")
}

func NewInventoryService(repo InventoryRepository, logger *slog.Logger) *InventoryService {
	return &InventoryService{Repo: repo, logger: logger}
}

func (s *InventoryService) GetAll() ([]models.Inventory, error) {
	return s.Repo.GetAll()
}

func (s *InventoryService) GetElementById(InventoryId int) (models.Inventory, error) {
	return s.Repo.GetElementById(InventoryId)
}

func (s *InventoryService) Delete(InventoryId int) error {
	//check that this id is valid
	return s.Repo.Delete(InventoryId)
}
