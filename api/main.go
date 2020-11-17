package main

import (
	"cardapio-virtual-api/src/config"
	"cardapio-virtual-api/src/router"
	"fmt"
	"log"
	"net/http"
)

const port int = 3333

func main() {

	config.Loader()
	router.Generate()

	fmt.Printf("Iniciando o server na porta %d\n", config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}
