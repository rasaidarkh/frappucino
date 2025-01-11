package models

import (
	"time"

	"github.com/lib/pq"
)

type Customer struct {
	CustomerID int            `json:"customer_id"`
	Name       string         `json:"name"`
	Age        int            `json:"age"`
	Sex        Sex            `json:"sex"`
	FirstOrder time.Time      `json:"first_order"`
	Allergens  pq.StringArray `json:"allergens"`
}

type Sex uint

const (
	Female Sex = iota
	Male
	Other
)
