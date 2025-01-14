package service

import (
	"context"
	"frappuccino/internal/models"
	"log/slog"
)

type InventoryRepository interface {
	GetAll(ctx context.Context) ([]models.Inventory, error)
	GetElementById(ctx context.Context, InventoryId int) (models.Inventory, error)
	Delete(ctx context.Context, InventoryId int) error
	Put(ctx context.Context, item models.Inventory) error
	Post(ctx context.Context, item models.Inventory) error
}

type InventoryService struct {
	Repo   InventoryRepository
	logger *slog.Logger
}

func (s *InventoryService) Put(ctx context.Context, item models.Inventory) error {
	//TODO implement me
	panic("implement me")
}

func (s *InventoryService) Post(ctx context.Context, item models.Inventory) error {
	//TODO implement me
	panic("implement me")
}

func NewInventoryService(repo InventoryRepository, logger *slog.Logger) *InventoryService {
	return &InventoryService{Repo: repo, logger: logger}
}

func (s *InventoryService) GetAll(ctx context.Context) ([]models.Inventory, error) {
	return s.Repo.GetAll(ctx)
}

func (s *InventoryService) GetElementById(ctx context.Context, InventoryId int) (models.Inventory, error) {
	return s.Repo.GetElementById(ctx, InventoryId)
}

func (s *InventoryService) Delete(ctx context.Context, InventoryId int) error {
	//check that this id is valid
	return s.Repo.Delete(ctx, InventoryId)
}
