package repositories

import (
	"cardapio-virtual-api/src/models"
	"database/sql"
	"errors"
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

// Find traz um cliente do banco de dados usando o CPF para realizar a busca
func (repository Customers) Find(doc string) (models.Customer, error) {
	var customer models.Customer

	stmt, err := repository.db.Prepare("SELECT id, document, e_table, update_at FROM customer WHERE document = ?")
	if err != nil {
		return models.Customer{}, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(doc).Scan(
		&customer.ID,
		&customer.Document,
		&customer.Table,
		&customer.UpdateAt,
	)
	if err != nil {
		return models.Customer{}, err
	}

	return customer, nil
}

// Update altera as informações de um cliente no banco de dados
func (repository Customers) Update(ID uint64, customer models.Customer) error {
	stmt, err := repository.db.Prepare("UPDATE customer SET document = ?, e_table = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(customer.Document, customer.Table, ID)
	if err != nil {
		return err
	}

	rowID, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowID == 0 {
		return errors.New("Resquesed item is not found")
	}

	return nil
}
