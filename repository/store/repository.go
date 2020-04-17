package store

import "github.com/dvasyanin/http-rest-api/models"

type ChatRepository interface {
	FindByRespond(idRespond int64) ([]*models.Chat, error)
}

type ClientsRepository interface {
	Create()
	FindByEmail()
}

type ResumeRepository interface {
	Create()
}

type VacancyRepository interface {
	Create()
}
