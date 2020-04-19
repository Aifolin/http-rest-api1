package pgstore

import (
	"context"
	"github.com/dvasyanin/http-rest-api/models"
)

type ClientRepository struct {
	store *Postgres
}

// Create client
func (r *ClientRepository) Create(e *models.Employers) error {
	query := `
		with 
			ins1 as (
				insert into clients (name, email, number_phone, active)
				VALUES ($1, $2, $3, $4) returning id
				),
			ins2 as (
				insert	into employers(id_client, about)
				VALUES ((select id from ins1), $5)
				)
			select id from ins1`

	return r.store.conn.QueryRow(context.Background(), query,
		e.Name, e.Email, e.NumberPhone, e.Active, e.About,
	).Scan(&e.ID)
}

// Delete client
func (r *ClientRepository) Delete(string) error {

	return nil
}

// Update client
func (r ClientRepository) Updated(*models.Employers) error {

	return nil
}

// Employers find by email
func (r *ClientRepository) FindByEmail(string) (*models.Employers, error) {

	return nil, nil
}
