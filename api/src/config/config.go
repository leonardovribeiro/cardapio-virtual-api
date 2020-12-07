package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringDbConnection é a string de conexão com o mysql
	StringDbConnection = ""

	// Port onde a api vai estar rodando
	Port = 0

	// SecretKey é a chave usada para assinar o token
	SecretKey []byte
)

// Loader vai inicializar as variáveis de ambiente
func Loader() {
	var err error

	err = godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		Port = 9000
	}

	StringDbConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
