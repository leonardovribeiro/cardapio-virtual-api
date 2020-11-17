package models

import (
	"time"
)

// Customer representa um cliente
type Customer struct {
	ID              uint64    `json:"id,omitempty"`
	Document        string    `json:"cpf,omitempty"`
	Table           uint8     `json:"table,omitempty"`
	FirstConnection time.Time `json:"first_connection,omitempty"`
	LastConnection  time.Time `json:"last_connection,omitempty"`
}
