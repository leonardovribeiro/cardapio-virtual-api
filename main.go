package main

import (
	"fmt"
	"log"
	"net/http"
)

const port int = 3333

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Bem vindo!")
	})

	fmt.Printf("Iniciando o server na porta %d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}
