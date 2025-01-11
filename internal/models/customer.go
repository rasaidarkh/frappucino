package models

import (
	"github.com/lib/pq"
	"time"
)

type Customer struct {
	CustomerID int            `json:"customer_id"`
	Name       string         `json:"name"`
	Age        int            `json:"age"`
	Sex        Sex            `json:"sex"`
	FirstOrder time.Time      `json:"first_order"`
	Allergens  pq.StringArray `json:"allergens"`
}

type Sex int

const (
	Female Sex = iota
	Male
	Other
)
