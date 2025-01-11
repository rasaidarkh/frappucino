package repository

import (
	"database/sql"
	"frappuccino/internal/models"
	"log/slog"
)

type orderRepositoryS struct {
	Db     *sql.DB
	Logger *slog.Logger
}

type OrderRepository interface {
	GetAll() ([]models.Menu, error)
}

func NewOrderRepository(db *sql.DB, logger *slog.Logger) *orderRepositoryS {
	return &orderRepositoryS{
		Db:     db,
		Logger: logger,
	}
}
