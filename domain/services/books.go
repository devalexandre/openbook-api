package services

import (
	"context"

	"github.com/devalexandre/openbook-api/adapters/database"
	"github.com/devalexandre/openbook-api/adapters/repo"
	"github.com/vingarcia/ksql"
)

var booksTable = ksql.NewTable("books", "id")

type Books struct {
	database.Database
}

func (s Books) List(ctx context.Context, offset, limit int) (total int, books []repo.Books, err error) {
	var countRow struct {
		Count int `ksql:"count"`
	}
	err = s.QueryOne(ctx, &countRow, "SELECT count(*) as count FROM books")
	if err != nil {
		return 0, nil, err
	}

	return countRow.Count, books, s.Query(ctx, &books, "SELECT * FROM books OFFSET ? LIMIT ?", offset, limit)
}
