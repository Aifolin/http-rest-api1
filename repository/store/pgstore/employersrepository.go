package pgstore

import (
	"github.com/dvasyanin/http-rest-api/models"
)

type ClientRepository struct {
	store *Postgres
}

// Create client
func (r *ClientRepository) Create(e *models.Employers) error {

	return nil
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
