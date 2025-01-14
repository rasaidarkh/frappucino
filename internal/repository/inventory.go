package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"frappuccino/internal/models"
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

func (r *InventoryRepository) logErrorAndReturn(err error, message string) error {
	r.Logger.Error(message, "error", err.Error())
	return err
}

func (r *InventoryRepository) Put(ctx context.Context, item models.Inventory) error {
	query := `UPDATE inventory SET inventory_name = $1, quantity = $2, unit = $3, allergens = $4 WHERE inventory_id = $5`
	stmt, err := r.Db.PrepareContext(ctx, query)
	if err != nil {
		return r.logErrorAndReturn(err, "failed to prepare update statement")
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, item.InventoryName, item.Quantity, item.Unit, item.Allergens, item.InventoryId)
	if err != nil {
		return r.logErrorAndReturn(err, "failed to execute update statement")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return r.logErrorAndReturn(err, "failed to get rows affected")
	}

	if rowsAffected == 0 {
		message := fmt.Sprintf("update failed, inventory item with ID %v does not exist", item.InventoryId)
		r.Logger.Warn(message)
		return errors.New(message)
	}

	r.Logger.Info("inventory item update was successful", "ID", item.InventoryId)
	return nil
}

func (r *InventoryRepository) Post(ctx context.Context, item models.Inventory) error {
	query := `INSERT INTO inventory (inventory_name, quantity, unit, allergens) VALUES ($1, $2, $3, $4) RETURNING inventory_id`
	stmt, err := r.Db.PrepareContext(ctx, query)
	if err != nil {
		return r.logErrorAndReturn(err, "failed to prepare insert statement")
	}
	defer stmt.Close()

	var lastId int
	err = stmt.QueryRowContext(ctx, item.InventoryName, item.Quantity, item.Unit, item.Allergens).Scan(&lastId)
	if err != nil {
		return r.logErrorAndReturn(err, "failed to execute insert statement")
	}

	r.Logger.Info("item was successfully inserted", "ID", lastId)
	return nil
}

func (r *InventoryRepository) GetAll(ctx context.Context) ([]models.Inventory, error) {
	query := `SELECT * FROM inventory`
	rows, err := r.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, r.logErrorAndReturn(err, "failed to execute select statement")
	}
	defer rows.Close()

	var inventoryItems []models.Inventory
	for rows.Next() {
		var inventory models.Inventory
		err := rows.Scan(&inventory.InventoryId, &inventory.InventoryName, &inventory.Quantity, &inventory.Unit, &inventory.Allergens)
		if err != nil {
			return nil, r.logErrorAndReturn(err, "failed to scan row")
		}
		inventoryItems = append(inventoryItems, inventory)
	}

	if err := rows.Err(); err != nil {
		return nil, r.logErrorAndReturn(err, "error iterating over rows")
	}

	r.Logger.Info("inventory items were transferred successfully")
	return inventoryItems, nil
}

func (r *InventoryRepository) GetElementById(ctx context.Context, inventoryId int) (models.Inventory, error) {
	query := `SELECT * FROM inventory WHERE inventory_id = $1`
	var inventory models.Inventory
	err := r.Db.QueryRowContext(ctx, query, inventoryId).
		Scan(&inventory.InventoryId, &inventory.InventoryName, &inventory.Quantity, &inventory.Unit, &inventory.Allergens)
	if err != nil {
		return models.Inventory{}, r.logErrorAndReturn(err, "failed to execute select by ID statement")
	}

	r.Logger.Info("inventory item was transferred successfully", "ID", inventory.InventoryId)
	return inventory, nil
}

func (r *InventoryRepository) Delete(ctx context.Context, inventoryId int) error {
	query := `DELETE FROM inventory WHERE inventory_id = $1`
	res, err := r.Db.ExecContext(ctx, query, inventoryId)
	if err != nil {
		return r.logErrorAndReturn(err, "failed to execute delete statement")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return r.logErrorAndReturn(err, "failed to get rows affected")
	}
	if rowsAffected == 0 {
		message := fmt.Sprintf("deletion was not successful, inventory item with ID %v does not exist", inventoryId)
		r.Logger.Warn(message)
		return errors.New(message)
	}

	r.Logger.Info("inventory item deletion was successful", "ID", inventoryId)
	return nil
}
