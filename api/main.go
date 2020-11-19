package main

import (
	"cardapio-virtual-api/src/config"
	"cardapio-virtual-api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	config.Loader()
	r := router.Generate()

	fmt.Printf("Iniciando o server na porta %d\n", config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r); err != nil {
		log.Fatal(err)
	}
}
