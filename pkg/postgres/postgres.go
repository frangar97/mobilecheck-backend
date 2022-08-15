package postgres

import (
	"database/sql"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func NewClient(connection string) (*sql.DB, error) {
	client, err := sql.Open("pgx", connection)

	if err != nil {
		return nil, err
	}

	err = client.Ping()

	if err != nil {
		return nil, err
	}

	return client, nil
}
