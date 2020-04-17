package pgstore

import (
	"github.com/dvasyanin/http-rest-api/repository/store"
	"github.com/jackc/pgx/v4"
)

type Postgres struct {
	conn              *pgx.Conn
	chatRepository    *ChatRepository
	clientRepository  *ClientRepository
	vacancyRepository *VacancyRepository
	resumeRepository  *ResumeRepository
}

// New ...
func New(db *pgx.Conn) store.Store {
	return &Postgres{
		conn: db,
	}
}

// Chat interface...
func (p *Postgres) Chat() store.ChatRepository {
	if p.chatRepository != nil {
		return p.chatRepository
	}

	p.chatRepository = &ChatRepository{
		store: p,
	}

	return p.chatRepository
}

// Client interface...
func (p *Postgres) Client() store.ClientsRepository {
	if p.clientRepository != nil {
		return p.clientRepository
	}

	p.clientRepository = &ClientRepository{
		store: p,
	}

	return p.clientRepository
}

// Vacancy interface...
func (p *Postgres) Vacancy() store.VacancyRepository {
	if p.vacancyRepository != nil {
		return p.vacancyRepository
	}

	p.vacancyRepository = &VacancyRepository{
		store: p,
	}

	return p.vacancyRepository
}

// Resume interface...
func (p *Postgres) Resume() store.ResumeRepository {
	if p.resumeRepository != nil {
		return p.resumeRepository
	}

	p.resumeRepository = &ResumeRepository{
		store: p,
	}

	return p.resumeRepository
}
