package repository

import (
	"database/sql"
	"frappuccino/internal/models"
	"log/slog"
)

type inventoryRepositoryS struct {
	Db     *sql.DB
	Logger *slog.Logger
}

func NewInventoryRepository(db *sql.DB, logger *slog.Logger) *inventoryRepositoryS {
	return &inventoryRepositoryS{
		Db:     db,
		Logger: logger,
	}
}

func (r *inventoryRepositoryS) GetAll() ([]models.Inventory, error) {
	stmt, err := r.Db.Prepare("SELECT * FROM inventory")
	if err != nil {
		r.Logger.Error(err.Error())
		return nil, err
	}

	r.Logger.Info("selec from inventory was successful")
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		r.Logger.Error(err.Error())
		return nil, err
	}
	var inventoryItems []models.Inventory

	for rows.Next() {
		var inventory models.Inventory
		if err = rows.Scan(&inventory.InventoryId, &inventory.InventoryName, &inventory.Quantity, &inventory.Unit, &inventory.Allergens); err != nil {
			r.Logger.Error(err.Error())
			return nil, err
		}
		inventoryItems = append(inventoryItems, inventory)
	}
	if rows.Err() != nil {
		r.Logger.Error(err.Error())
		return nil, rows.Err()
	}
	r.Logger.Info("inventory items were transferred successfully")
	return inventoryItems, nil
}
