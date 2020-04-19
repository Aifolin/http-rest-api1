package pgstore

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dvasyanin/http-rest-api/models"
)

type ChatRepository struct {
	store *Postgres
}

// CreateMessage message
func (r *ChatRepository) Create(chat *models.Chat) error {
	query := `insert into chat 
				(message, id_respond, id_client) 
				VALUES ($1, $2, $3) returning id_message`

	return r.store.conn.QueryRow(context.Background(), query,
		chat.Message, chat.RespondID, chat.ClientID,
	).Scan(&chat.MessageID)
}

// Get chat as part of the response
func (r *ChatRepository) FindByRespond(id int64) ([]*models.ChatByRespond, error) {
	var chat []*models.ChatByRespond

	query := `
		select c2.message,
			   c2.created,
			   cl.name
		from respond r
				left join resume r2 on r.id_resume = r2.id
				left join vacancies v on r.id_vacancy = v.id
				left join clients c on r2.id_client = c.id and v.id_client = c.id
				left join chat c2 on r.id = c2.id_respond
				left join clients cl on c2.id_client = cl.id
		where r.id = $1
		order by c2.created`

	row, err := r.store.conn.Query(context.Background(), query, id)

	if err != nil {
		return nil, err
	}

	var msg, name sql.NullString
	var created sql.NullTime

	for row.Next() {
		err := row.Scan(&msg, &created, &name)
		if err != nil {
			return nil, err
		}

		chat = append(chat, &models.ChatByRespond{
			Message: msg.String,
			Created: created.Time,
			Name:    name.String,
		})
	}

	if len(chat) == 0 {
		return nil, errors.New("chat not found")
	}

	return chat, nil
}
