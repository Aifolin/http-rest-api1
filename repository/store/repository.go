package store

import "github.com/dvasyanin/http-rest-api/models"

type ChatRepository interface {
	Create(*models.Chat) error
	FindByRespond(int64) ([]*models.ChatByRespond, error)
}

type ClientsRepository interface {
	Create(*models.Employers) error
	Delete(string) error
	Updated(*models.Employers) error
	FindByEmail(string) (*models.Employers, error)
}

type ResumeRepository interface {
	Create()
}

type VacancyRepository interface {
	Create()
}
