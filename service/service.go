package service

import (
	"github.com/dvasyanin/http-rest-api/models"
	"github.com/dvasyanin/http-rest-api/repository/store"
)

type Service interface {
	GetChatByRespond(int64) ([]*models.ChatByRespond, error)
	CreateMessage(*models.Chat) error
	CreateEmployer(*models.Employers) error
	DeleteEmployer(string) error
	EmployerFindByEmail(*models.Employers) error
	UpsertEmployer(*models.Employers) error
}

type service struct {
	store store.Store
}

func NewService(store store.Store) Service {
	return &service{
		store: store,
	}
}

func (s *service) GetChatByRespond(id int64) ([]*models.ChatByRespond, error) {
	// default count = 100
	// default page = 1
	// https://api.slack.com/methods/team.accessLogs
	chat, err := s.store.Chat().FindByRespond(id)
	if err != nil {
		return nil, err
	}

	return chat, nil
}

func (s *service) CreateMessage(chat *models.Chat) error {
	if err := s.store.Chat().Create(chat); err != nil {
		return err
	}

	return nil
}

func (s *service) CreateEmployer(employer *models.Employers) error {
	if err := s.store.Client().CreateEmployer(employer); err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteEmployer(email string) error {
	if err := s.store.Client().DeleteEmployer(email); err != nil {
		return err
	}

	return nil
}

func (s *service) EmployerFindByEmail(employer *models.Employers) error {
	if err := s.store.Client().EmployerFindByEmail(employer); err != nil {
		return err
	}

	return nil
}

func (s *service) UpsertEmployer(employer *models.Employers) error {
	if err := s.store.Client().UpsertEmployer(employer); err != nil {
		return err
	}

	return nil
}
