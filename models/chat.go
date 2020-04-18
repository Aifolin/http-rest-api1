package models

import "time"

type Chat struct {
	MessageID int64 `sql:"id_message"`
	Created   time.Time
	Updated   time.Time
	Message   string
	RespondID int64 `sql:"id_respond"`
	ClientID  int64 `sql:"id_client"`
}

type ChatByRespond struct {
	Message string    `json:"message"`
	Created time.Time `json:"created"`
	Name    string    `json:"sender"`
}
