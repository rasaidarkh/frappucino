package repository

import (
	"database/sql"
	"frappuccino/internal/models"
	"log/slog"
)

type menuRepositoryS struct {
	Db     *sql.DB
	Logger *slog.Logger
}

type MenuRepository interface {
	GetAll() ([]models.Menu, error)
}

func NewMenuRepository(db *sql.DB, logger *slog.Logger) *menuRepositoryS {
	return &menuRepositoryS{
		Db:     db,
		Logger: logger,
	}
}
