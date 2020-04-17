package store

type Store interface {
	Chat() ChatRepository
	Client() ClientsRepository
	Vacancy() VacancyRepository
	Resume() ResumeRepository
}
