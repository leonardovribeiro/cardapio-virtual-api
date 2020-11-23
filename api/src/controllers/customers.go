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
	"strconv"

	"github.com/gorilla/mux"
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

	err = customer.Prepare("login")
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
	param := mux.Vars(r)

	var customer models.Customer
	customer.Document = param["customerDoc"]

	err := customer.Prepare("find")

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewCustomersRepository(db)
	customer, err = repository.Find(customer.Document)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, customer)

}

// UpdateCustomer atualiza um cliente no banco
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerID, err := strconv.ParseUint(params["customerID"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var customer models.Customer
	err = json.Unmarshal(reqBody, &customer)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	err = customer.Prepare("update")

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
	err = repository.Update(customerID, customer)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

// DeleteCustomer deleta um cliente no banco
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Deleta um cliente")))
}
