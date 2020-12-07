package repositories

import (
	"cardapio-virtual-api/src/models"
	"database/sql"
	"fmt"
	"time"
)

// Products representa um repositório de produtos
type Products struct {
	db *sql.DB
}

// NewProductsRepository cria um novo repositório de produtos
func NewProductsRepository(db *sql.DB) *Products {
	return &Products{db}
}

// Create cria um novo produto no banco de dados
func (repository Products) Create(product models.Product) (uint64, error) {

	stmt, err := repository.db.Prepare("INSERT INTO products(name, description, price, created_at) VALUES(?,?,?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	createdAt := time.Now()

	res, err := stmt.Exec(product.Name, product.Description, product.Price, createdAt)
	if err != nil {
		return 0, err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return uint64(lastID), nil
}

// Fetch busca todos os produtos no banco de dados
func (repository Products) Fetch() ([]models.Product, error) {
	var products []models.Product

	rows, err := repository.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt)
		if err != nil {
			return []models.Product{}, err
		}
		products = append(products, product)
	}

	return products, nil
}

// GetByID busca um produto no banco de dados usando o ID
func (repository Products) GetByID(ID uint64) (models.Product, error) {
	row, err := repository.db.Query("SELECT * FROM products WHERE id=?", ID)
	if err != nil {
		return models.Product{}, err
	}
	defer row.Close()

	var product models.Product

	if row.Next() {
		err = row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt)
		if err != nil {
			return models.Product{}, err
		}
	}

	return product, nil
}

// Update cria um novo produto no banco de dados
func (repository Products) Update(ID uint64, product models.Product) error {
	stmt, err := repository.db.Prepare("UPDATE products SET name=?, description=?, price=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	fmt.Println(product)
	_, err = stmt.Exec(product.Name, product.Description, product.Price, ID)
	if err != nil {
		return err
	}

	return nil
}

// Delete cria um novo produto no banco de dados
func (repository Products) Delete(ID uint64) error {
	stmt, err := repository.db.Prepare("DELETE FROM products WHERE id=?")
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
