package repository

import (
	"github.com/BrunoCesarAngst/golang/adapter/repository/fixture"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestTransactionRepositoryDbInsert(t *testing.T) {
	migrateDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrateDir)
	defer fixture.Down(db, migrateDir)

	repository := NewTransactionRepositoryDb(db)
	err := repository.Insert("1", "1", 2, "Approved", "")
	assert.Nil(t, err)
}
