package db

import (
	"database/sql"
	"errors"
	"fmt"
)

func CreateConnection(url string) (*sql.DB, error) {
	db, err := sql.Open("libsql", url)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to open db %s: %s\n", url, err))
	}
	defer db.Close()
	return db, nil
}
