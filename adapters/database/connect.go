package database

import (
	"context"
	"fmt"
	"os"

	"github.com/vingarcia/ksql"
	"github.com/vingarcia/ksql/adapters/kpgx"
)

func NewConnect() (*ksql.DB, error) {

	db, err := kpgx.New(context.Background(), os.Getenv("DATABASEURL"), ksql.Config{})

	if err != nil {
		return nil, fmt.Errorf("unable to connect to database, %v", err)
	}

	return &db, nil
}
