package models

import "time"

type Chat struct {
	TableName struct{}  `pg:"chat" json:"-"`
	Message   string    `json:"message"`
	Created   time.Time `json:"created"`
	Name      string    `json:"sender"`
}
