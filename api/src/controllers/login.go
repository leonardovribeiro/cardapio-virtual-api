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

// Login é responsável por autenticar um cliente na API
func Login(w http.ResponseWriter, r *http.Request) {

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

	customer.Prepare("find")

	db, err := database.Connection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewCustomersRepository(db)
	customerFound, err := repository.GetByDoc(customer.Document)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if customerFound.ID != 0 {
		err = repository.Update(customerFound.ID, customer)
		if err != nil {
			responses.Error(w, http.StatusInternalServerError, err)
			return
		}

	} else {
		customerFound.ID, err = repository.Create(customer)
		if err != nil {
			responses.Error(w, http.StatusInternalServerError, err)
			return
		}
	}

	token, _ := authentication.CreateToken(customerFound.ID)

	responses.JSON(w, http.StatusOK, token)
}
