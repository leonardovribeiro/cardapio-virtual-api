package controllers

import (
	"cardapio-virtual-api/src/database"
	"cardapio-virtual-api/src/models"
	"cardapio-virtual-api/src/repositories"
	"cardapio-virtual-api/src/responses"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// CreateCustomer insere um cliente no banco de dados
func CreateCustomer(w http.ResponseWriter, r *http.Request) {

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var customer models.Customer
	err = json.Unmarshal(requestBody, &customer)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	err = customer.Prepare()
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewCustomersRepository(db)
	customerID, err := repository.Create(customer)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, customerID)

}

// FindAllCustomer busca todos os clientes no banco
func FindAllCustomer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Busca todos os clientes")))
}

// FindCustomer busca um cliente no banco
func FindCustomer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Busca um cliente")))
}

// UpdateCustomer atualiza um cliente no banco
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Atualiza um cliente")))
}

// DeleteCustomer deleta um cliente no banco
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Deleta um cliente")))
}
