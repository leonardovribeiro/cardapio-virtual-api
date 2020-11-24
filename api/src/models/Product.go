package models

// Product representa um produto no banco de dados
type Product struct {
	ID          uint64   `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Price       Currency `json:"price,omitempty"`
}
