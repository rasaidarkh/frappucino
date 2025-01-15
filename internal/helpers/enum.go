package helpers

import "frappuccino/internal/models"

func IsValid(t models.TransactionType) bool {
	switch t {
	case 0, 1, 2:
		return true
	default:
		return false
	}
}
