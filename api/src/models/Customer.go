package models

import (
	"errors"
	"strings"
	"time"
)

// Customer representa um cliente
type Customer struct {
	ID       uint64    `json:"id,omitempty"`
	Document string    `json:"cpf,omitempty"`
	Table    uint8     `json:"table,omitempty"`
	UpdateAt time.Time `json:"update_at,omitempty"`
}

// Prepare vai chamar os métodos para validar e formatar o cliente recebido
func (customer *Customer) Prepare() error {
	err := customer.validator()
	if err != nil {
		return err
	}

	customer.format()

	return nil
}

func (customer *Customer) validator() error {
	if customer.Document == "" {
		return errors.New("O número do documento é obrigatório e não pode estar em branco")
	}

	if customer.Table == 0 {
		return errors.New("O número da mesa é obrigatório e não pode estar em branco")
	}

	return nil
}

func (customer *Customer) format() {
	customer.Document = strings.TrimSpace(customer.Document)
	customer.Document = strings.ReplaceAll(customer.Document, ".", "")
	customer.Document = strings.ReplaceAll(customer.Document, "-", "")
}
