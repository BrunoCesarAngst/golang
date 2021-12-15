package entity

import "errors"

type Transaction struct {
	ID           string
	AccountID    string
	Amount       float64
	Status       string
	ErrorMassage string
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (t *Transaction) IsValid() error {
	if t.Amount > 1000 {
		return errors.New("Passou do limite")
	}
	if t.Amount < 1 {
		return errors.New("Abaixo do valor permitido")
	}
	return nil
}
