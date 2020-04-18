package pgstore_test

import (
	"github.com/dvasyanin/http-rest-api/models"
	"github.com/dvasyanin/http-rest-api/repository/store/pgstore"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChatRepository_Create(t *testing.T) {
	db, teardown := pgstore.TestDb(t, databaseURL)
	defer teardown("chat")

	s := pgstore.New(db)
	ch := models.TestChat(t)
	assert.NoError(t, s.Chat().Create(ch))
	assert.NotZero(t, ch.MessageID)
}

func TestChatRepository_FindByRespond(t *testing.T) {
	db, teardown := pgstore.TestDb(t, databaseURL)
	defer teardown("chat")

	s := pgstore.New(db)
	ch := models.TestChat(t)

	chat, err := s.Chat().FindByRespond(ch.MessageID)
	assert.Error(t, err)
	assert.Nil(t, chat)

	err = s.Chat().Create(ch)
	assert.NoError(t, err)
	assert.NotNil(t, ch)
}
