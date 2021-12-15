package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransactionWithAmountGreaterThan1000(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 2000

	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "Passou do limite", err.Error())
}

func TestTransactionWithAmountLessThan1(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "1"
	transaction.AccountID = "1"
	transaction.Amount = 0

	err := transaction.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "Abaixo do valor permitido", err.Error())
}
