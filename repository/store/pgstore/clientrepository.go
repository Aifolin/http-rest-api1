package pgstore

import (
	"context"
	"github.com/dvasyanin/http-rest-api/models"
)

type ClientRepository struct {
	store *Postgres
}

// CreateMessage client
func (r *ClientRepository) CreateEmployer(e *models.Employers) error {
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

// Delete employer
func (r *ClientRepository) DeleteEmployer(string) error {

	return nil
}

// Employer find by email
func (r *ClientRepository) EmployerFindByEmail(*models.Employers) error {

	return nil
}

// Upsert employer
func (r *ClientRepository) UpsertEmployer(*models.Employers) error {

	return nil
}
