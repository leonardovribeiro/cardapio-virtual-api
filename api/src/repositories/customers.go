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

	stmt, err := repository.db.Prepare("INSERT INTO customers(doc, table_c) VALUES(?,?)")
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

// Fetch busca todos os clientes no banco de dados
func (repository Customers) Fetch() ([]models.Customer, error) {

	rows, err := repository.db.Query("SELECT * FROM customers")
	if err != nil {
		return []models.Customer{}, err
	}
	defer rows.Close()

	var customers []models.Customer

	for rows.Next() {
		var customer models.Customer
		err = rows.Scan(&customer.ID, &customer.Document, &customer.Table, &customer.UpdateAt)
		if err != nil {
			return []models.Customer{}, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

// GetByID traz um cliente do banco de dados usando o id para realizar a busca
func (repository Customers) GetByID(ID uint64) (models.Customer, error) {

	var customer models.Customer

	stmt, err := repository.db.Prepare("SELECT id, doc, table_c, update_at FROM customers WHERE id = ?")
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

// GetByDoc traz um cliente do banco de dados usando o documento para realizar a busca
func (repository Customers) GetByDoc(doc string) (models.Customer, error) {

	var customer models.Customer

	stmt, err := repository.db.Prepare("SELECT id, doc, table_c, update_at FROM customer WHERE document = ?")
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
	stmt, err := repository.db.Prepare("UPDATE customers SET doc = ?, table_c = ? where id = ?")
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

// Delete altera as informações de um cliente no banco de dados
func (repository Customers) Delete(ID uint64) error {
	stmt, err := repository.db.Prepare("DELETE FROM customers WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(ID)
	if err != nil {
		return err
	}
	return nil
}
