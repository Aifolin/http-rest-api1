package models

import (
	"testing"
)

func TestChat(t *testing.T) *Chat {
	t.Helper()

	return &Chat{
		Message:   "test message",
		RespondID: 1,
		ClientID:  1,
	}
}

func TestEmployer(t *testing.T) *Employers {
	t.Helper()

	return &Employers{
		Name:        "test name",
		Email:       "test email",
		NumberPhone: "test phone",
		Active:      true,
		About:       "I`am employers",
	}
}
