package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"frappuccino/internal/models"

	"log/slog"
)

type InventoryRepository struct {
	Db     *sql.DB
	Logger *slog.Logger
}

func NewInventoryRepository(db *sql.DB, logger *slog.Logger) *InventoryRepository {
	return &InventoryRepository{
		Db:     db,
		Logger: logger,
	}
}

func (r *InventoryRepository) Put(item models.Inventory) error {
	//TODO implement me
	panic("implement me")
}

func (r *InventoryRepository) Post(item models.Inventory) error {
	stmt, err := r.Db.Prepare("INSERT INTO inventory (inventory_name,quantity,unit,allergens) VALUES($1,$2,$3,$4) RETURNING inventory_id")
	if err != nil {
		r.Logger.Error(err.Error())
		return err
	}

	var LastId int
	err = stmt.QueryRow(item.InventoryName, item.Quantity, item.Unit, item.Allergens).Scan(LastId)
	if err != nil {
		r.Logger.Error(err.Error())
		return err
	}
	r.Logger.Info("item was successfully inserted", "ID", LastId)
	return nil
}

func (r *InventoryRepository) GetAll() ([]models.Inventory, error) {
	stmt, err := r.Db.Prepare("SELECT * FROM inventory")
	if err != nil {
		r.Logger.Error(err.Error())
		return nil, err
	}

	r.Logger.Info("inventory select preparation was successful")
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		r.Logger.Error(err.Error())
		return nil, err
	}
	var inventoryItems []models.Inventory

	for rows.Next() {
		var inventory models.Inventory
		if err = rows.Scan(
			&inventory.InventoryId, &inventory.InventoryName,
			&inventory.Quantity, &inventory.Unit, &inventory.Allergens); err != nil {
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

func (r *InventoryRepository) GetElementById(InventoryId int) (models.Inventory, error) {
	stmt, err := r.Db.Prepare("SELECT * FROM inventory WHERE inventory_id = $1")

	if err != nil {
		r.Logger.Error(err.Error())
		return models.Inventory{}, err
	}

	r.Logger.Info("inventory select preparation was successful")
	defer stmt.Close()
	var inventoryItem models.Inventory
	if err = stmt.QueryRow(InventoryId).Scan(
		&inventoryItem.InventoryId, &inventoryItem.InventoryName,
		&inventoryItem.Quantity, &inventoryItem.Unit, &inventoryItem.Allergens); err != nil {
		r.Logger.Error(err.Error())
		return models.Inventory{}, err
	}

	r.Logger.Info("inventory item was transferred successfully")
	return inventoryItem, nil
}

func (r *InventoryRepository) Delete(InventoryId int) error {
	// const op = "repository.inventory.Delete"
	stmt, err := r.Db.Prepare("DELETE  FROM inventory WHERE inventory_id = $1")
	if err != nil {
		r.Logger.Error(err.Error())
		return err
	}
	res, err := stmt.Exec(InventoryId)
	if err != nil {
		r.Logger.Error(err.Error())
		return err
	}
	n, err := res.RowsAffected()
	if err == nil {
		r.Logger.Error(fmt.Sprint(err))
		return err
	}
	if n == 0 {
		val := fmt.Sprintf("deletion was not successful, %v does not exist", InventoryId)
		r.Logger.Warn(val)
		return errors.New(val)

	}
	r.Logger.Info("inventory item deletion was successful")

	return nil
}
