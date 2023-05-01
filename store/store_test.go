package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M){
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "root:12345678@tcp(localhost:3306)/restapi_test"
	}

	os.Exit(m.Run())
}