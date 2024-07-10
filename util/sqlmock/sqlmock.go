package sqlmock

import (
	"database/sql"
	"log"

	gosqlmock "github.com/DATA-DOG/go-sqlmock"
)

func NewMock() (*sql.DB, gosqlmock.Sqlmock) {
	db, mock, err := gosqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}
