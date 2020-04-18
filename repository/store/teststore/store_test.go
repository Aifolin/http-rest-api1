package pgstore_test

import (
	"github.com/subosito/gotenv"
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	_ = gotenv.Load(".env")
	databaseURL = os.Getenv("DATABASE_URL_TEST")
	if databaseURL == "" {
		databaseURL = "host=localhost port=5434 dbname=test user=test password=test123 sslmode=disable"
	}

	os.Exit(m.Run())
}
