package tests

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func NewDBMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	return db, mock, nil
}
