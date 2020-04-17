package pgstore

import (
	"context"
	"errors"
	"github.com/dvasyanin/http-rest-api/models"
)

type ChatRepository struct {
	store *Postgres
}

// Get chat as part of the response
func (r *ChatRepository) FindByRespond(id int64) ([]*models.Chat, error) {
	var chat []*models.Chat

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

	for row.Next() {
		var ch models.Chat
		err := row.Scan(&ch)
		if err != nil {
			return nil, err
		}

		chat = append(chat, &ch)
	}

	if len(chat) == 0 {
		return nil, errors.New("chat not found")
	}

	return chat, nil
}
