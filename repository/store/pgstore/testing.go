package pgstore

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"strings"
	"testing"
)

// TestDB ...
func TestDb(t *testing.T, databaseURL string) (*pgx.Conn, func(...string)) {
	t.Helper()

	db, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(context.Background()); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			_, err := db.Exec(context.Background(), fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
			if err != nil {
				t.Fatal(err)
			}
		}

		if err := db.Close(context.Background()); err != nil {
			t.Fatal(err)
		}
	}
}
