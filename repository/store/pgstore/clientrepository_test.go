package pgstore_test

import (
	"github.com/dvasyanin/http-rest-api/models"
	"github.com/dvasyanin/http-rest-api/repository/store/pgstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClientRepository_CreateEmployers(t *testing.T) {
	db, teardown := pgstore.TestDb(t, databaseURL)
	defer teardown("clients", "employers")

	s := pgstore.New(db)
	e := models.TestEmployer(t)

	assert.NoError(t, s.Client().CreateEmployer(e))
	assert.NotZero(t, e.ID)
}
