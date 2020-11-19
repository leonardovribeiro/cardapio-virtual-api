package repositories

import (
	"cardapio-virtual-api/src/models"
	"database/sql"
)

// Customers representa um repositório de clientes
type Customers struct {
	db *sql.DB
}

// NewCustomersRepository cria um novo repositório de clientes
func NewCustomersRepository(db *sql.DB) *Customers {
	return &Customers{db}
}

// Create insere um novo cliente no banco de dados
func (repository Customers) Create(customer models.Customer) (uint64, error) {
	stmt, err := repository.db.Prepare("INSERT INTO customer(document, e_table) VALUES(?,?);")

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(customer.Document, customer.Table)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}
