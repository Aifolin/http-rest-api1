package pgstore

import (
	"context"
	"fmt"
	"github.com/dvasyanin/http-rest-api/repository/store"
	"github.com/jackc/pgx/v4"
	"os"
)

type Postgres struct {
	conn              *pgx.Conn
	chatRepository    *ChatRepository
	clientRepository  *ClientRepository
	vacancyRepository *VacancyRepository
	resumeRepository  *ResumeRepository
}

// New ...
func New(db *pgx.Conn) *Postgres {
	return &Postgres{
		conn: db,
	}
}

// Conn ...
func (p *Postgres) db() *pgx.Conn {
	if p != nil {
		return p.conn
	}

	return nil
}

// Open connect...
func (p *Postgres) Connect(dbURL string) {

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	p.conn = conn

	return
}

// Close connect...
func (p *Postgres) Close() {
	if p != nil {
		p.conn.Close(context.Background())
	}

	return
}

// Ping data base...
func (p Postgres) Ping() error {
	_, err := p.conn.Exec(context.Background(), "SELECT 1")
	if err != nil {
		return err
	}

	return nil
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
