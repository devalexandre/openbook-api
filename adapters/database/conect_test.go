package database

import (
	"os"
	"testing"
)

func TestConnect(t *testing.T) {

	os.Setenv("DATABASEURL", "postgresql://openbook:4gh8fkKrgGl792_lRjk8vA@free-tier14.aws-us-east-1.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full&options=--cluster=openbook-5562")

	t.Run("TestConnect", func(t *testing.T) {
		_, err := NewConnect()
		if err != nil {
			t.Errorf("Error on connect to database, %v", err)
		}
	})
}
