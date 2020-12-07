package repositories

import (
	"cardapio-virtual-api/src/models"
	"database/sql"
	"time"
)

// Orders representa um repositório de pedidos
type Orders struct {
	db *sql.DB
}

// NewOrdersRepository cria um novo repositório de pedidos
func NewOrdersRepository(db *sql.DB) *Orders {
	return &Orders{db}
}

// Create insere um pedido no banco de dados
func (repo Orders) Create(order models.Order) (uint64, error) {

	tx, err := repo.db.Begin()
	if err != nil {
		return 0, err
	}
	stmt, err := tx.Prepare("INSERT INTO orders(customer_id, comments,created_at) VALUES(?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(order.Customer.ID, order.Comments, order.Total, time.Now().Format("2020-10-10"))
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	orderID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	// Laço que percorre todos os produtos do pedido e os insere no banco
	for _, product := range order.Products {
		stmt, err = tx.Prepare("INSERT INTO order_items(order_id, product_id, quantity) VALUES(?, ?, ?)")
		if err != nil {
			return 0, err
		}
		defer stmt.Close()

		_, err := stmt.Exec(orderID, product.ID, product.Quantity)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return uint64(orderID), nil
}
