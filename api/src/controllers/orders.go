package controllers

import (
	"cardapio-virtual-api/src/authentication"
	"cardapio-virtual-api/src/database"
	"cardapio-virtual-api/src/models"
	"cardapio-virtual-api/src/repositories"
	"cardapio-virtual-api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// CreateOrder cria uma novo pedido
func CreateOrder(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	var order models.Order
	err = json.Unmarshal(reqBody, &order)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Extrai o usuário que está no cabeçalho da requisição e o atribui ao usuário que fez o pedido
	order.Customer.ID, err = authentication.ExtractUserID(r)
	if err != nil {
		responses.Error(w, http.StatusForbidden, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	defer db.Close()

	repository := repositories.NewOrdersRepository(db)
	orderID, err := repository.Create(order)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusCreated, orderID)

}

// FetchOrder busca todos os pedidos no banco de dados
func FetchOrder(w http.ResponseWriter, r *http.Request) {}

// GetOrderByID busca um pedido usando o ID
func GetOrderByID(w http.ResponseWriter, r *http.Request) {}

// UpdateOrder atualzia um pedido no banco de dados
func UpdateOrder(w http.ResponseWriter, r *http.Request) {}

// DeleteOrder exclui um pedido no banco de dados
func DeleteOrder(w http.ResponseWriter, r *http.Request) {}
