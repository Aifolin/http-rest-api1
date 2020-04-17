package service

import (
	"github.com/dvasyanin/http-rest-api/models"
	"github.com/dvasyanin/http-rest-api/repository/store/pgstore"
)

type Service interface {
	// default count = 100
	// default page = 1
	// https://api.slack.com/methods/team.accessLogs
	GetChatByRespond(id int64) ([]*models.Chat, error)
}

type service struct {
	store *pgstore.Postgres
}

func NewService(store *pgstore.Postgres) Service {
	return &service{
		store: store,
	}
}

func (s *service) GetChatByRespond(id int64) ([]*models.Chat, error) {
	s.store.Connect("")
	defer s.store.Close()

	chat, err := s.store.Chat().FindByRespond(id)
	if err != nil {
		return nil, err
	}

	return chat, nil
}
