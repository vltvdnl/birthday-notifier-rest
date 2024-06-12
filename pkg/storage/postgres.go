package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	*sql.DB
}

func New(url string) (*Storage, error) {
	const log_op = "Storage.New"
	db, err := sql.Open("postgres", url)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", log_op, err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", log_op, err)
	}

	return &Storage{db}, nil
}

//TODO: methods to search and others things....
