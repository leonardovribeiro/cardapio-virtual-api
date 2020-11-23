package repositories

import (
	"cardapio-virtual-api/src/models"
	"database/sql"
	"fmt"
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

	fmt.Println(customer)
	stmt, err := repository.db.Prepare("INSERT INTO customer(document, e_table) VALUES(?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

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

// Find traz um cliente do banco de dados usando o id para realizar a busca
func (repository Customers) Find(ID uint64) (models.Customer, error) {

	var customer models.Customer

	stmt, err := repository.db.Prepare("SELECT id, document, e_table, update_at FROM customer WHERE id = ?")
	if err != nil {
		return models.Customer{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(ID)
	if err != nil {
		return models.Customer{}, err
	}

	if rows.Next() {
		err = rows.Scan(&customer.ID, &customer.Document, &customer.Table, &customer.UpdateAt)
		if err != nil {
			return models.Customer{}, err
		}
	}

	return customer, nil
}

// FindByDoc traz um cliente do banco de dados usando o id para realizar a busca
func (repository Customers) FindByDoc(doc string) (models.Customer, error) {

	var customer models.Customer

	stmt, err := repository.db.Prepare("SELECT id, document, e_table, update_at FROM customer WHERE document = ?")
	if err != nil {
		return models.Customer{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(doc)
	if err != nil {
		return models.Customer{}, err
	}

	if rows.Next() {
		err = rows.Scan(&customer.ID, &customer.Document, &customer.Table, &customer.UpdateAt)
		if err != nil {
			return models.Customer{}, err
		}
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

	_, err = stmt.Exec(customer.Document, customer.Table, ID)
	if err != nil {
		return err
	}

	return nil
}
