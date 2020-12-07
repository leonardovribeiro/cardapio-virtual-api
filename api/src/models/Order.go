package models

import (
	"time"
)

// Order representa um pedido
type Order struct {
	ID        uint64    `json:"id"`
	Customer  uint64    `json:"customer_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}
