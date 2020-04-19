package models

import "time"

type Employers struct {
	ID          int64
	Created     time.Time
	Updated     time.Time
	Name        string
	Email       string
	NumberPhone string
	Active      bool
	About       string
}
