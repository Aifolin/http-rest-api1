package store

import "github.com/dvasyanin/http-rest-api/models"

type ChatRepository interface {
	Create(*models.Chat) error
	FindByRespond(int64) ([]*models.ChatByRespond, error)
}

type ClientsRepository interface {
	CreateEmployer(*models.Employers) error
	DeleteEmployer(string) error
	EmployerFindByEmail(*models.Employers) error
	UpsertEmployer(*models.Employers) error
}

type ResumeRepository interface {
	Create()
}

type VacancyRepository interface {
	Create()
}
