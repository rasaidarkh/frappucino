package service

import (
	"context"
	"errors"
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

func NewInventoryService(repo InventoryRepository, logger *slog.Logger) *InventoryService {
	return &InventoryService{Repo: repo, logger: logger}
}

func (s *InventoryService) GetAll(ctx context.Context) ([]models.Inventory, error) {
	items, err := s.Repo.GetAll(ctx)
	if err != nil {
		s.logger.Error("GetAll() for inventory failed", slog.String("error", err.Error()))
		return nil, err
	}
	return items, nil
}

func (s *InventoryService) GetElementById(ctx context.Context, InventoryId int) (models.Inventory, error) {
	item, err := s.Repo.GetElementById(ctx, InventoryId)
	if err != nil {
		s.logger.Error("Failed to fetch inventory item by ID", slog.Int("ID", InventoryId), slog.String("error", err.Error()))
		return models.Inventory{}, err
	}
	return item, nil
}

func (s *InventoryService) Delete(ctx context.Context, InventoryId int) error {
	err := s.Repo.Delete(ctx, InventoryId)
	if err != nil {
		s.logger.Error("Failed to delete inventory item", slog.Int("ID", InventoryId), slog.String("error", err.Error()))
		return err
	}
	s.logger.Info("Successfully deleted inventory item", slog.Int("ID", InventoryId))
	return nil
}

func (s *InventoryService) Put(ctx context.Context, item models.Inventory) error {
	if item.InventoryId <= 0 {
		err := errors.New("invalid item ID")
		s.logger.Error("Failed to update inventory item", slog.String("error", err.Error()))
		return err
	}

	err := s.Repo.Put(ctx, item)
	if err != nil {
		s.logger.Error("Failed to update inventory item", slog.Int("ID", item.InventoryId), slog.String("error", err.Error()))
		return err
	}

	s.logger.Info("Successfully updated inventory item", slog.Int("ID", item.InventoryId))
	return nil
}

func (s *InventoryService) Post(ctx context.Context, item models.Inventory) error {
	if item.InventoryName == "" || item.Quantity < 0 {
		err := errors.New("invalid item data")
		s.logger.Error("Failed to create inventory item", slog.String("error", err.Error()))
		return err
	}

	err := s.Repo.Post(ctx, item)
	if err != nil {
		s.logger.Error("Failed to create inventory item", slog.String("error", err.Error()))
		return err
	}

	s.logger.Info("Successfully created inventory item", slog.String("InventoryName", item.InventoryName))
	return nil
}
