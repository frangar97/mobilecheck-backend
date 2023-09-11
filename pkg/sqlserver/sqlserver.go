package sqlserver

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

func NewClient(connection string) (*sql.DB, error) {
	client, err := sql.Open("mssql", connection)

	if err != nil {
		return nil, err
	}

	err = client.Ping()

	if err != nil {
		return nil, err
	}

	return client, nil
}
