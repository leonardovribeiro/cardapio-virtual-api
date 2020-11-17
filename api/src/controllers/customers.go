package controllers

import (
	"cardapio-virtual-api/src/database"
	"cardapio-virtual-api/src/models"
	"cardapio-virtual-api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CreateCustomer insere um cliente no banco de dados
func CreateCustomer(w http.ResponseWriter, r *http.Request) {

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var customer models.Customer
	err = json.Unmarshal(requestBody, &customer)
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Connection()
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewCustomersRepository(db)
	customerID, err := repository.Create(customer)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("ID inserido: %d", customerID)))

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
