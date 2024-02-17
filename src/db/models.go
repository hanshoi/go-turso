// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql"
)

type Person struct {
	ID        int64
	FirstName string
	LastName  string
	Title     sql.NullString
	Company   sql.NullString
	Age       sql.NullInt64
	Married   bool
}

type SearchablePerson struct {
	FirstName string
	LastName  string
	Title     string
	Company   string
}