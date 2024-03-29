// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
)

const countPeople = `-- name: CountPeople :one
SELECT COUNT(id) FROM people
`

func (q *Queries) CountPeople(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countPeople)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createPerson = `-- name: CreatePerson :exec
INSERT INTO people (first_name, last_name, title, company, age, married) VALUES (?, ?, ?, ?, ?, ?)
`

type CreatePersonParams struct {
	FirstName string
	LastName  string
	Title     sql.NullString
	Company   sql.NullString
	Age       sql.NullInt64
	Married   bool
}

func (q *Queries) CreatePerson(ctx context.Context, arg CreatePersonParams) error {
	_, err := q.db.ExecContext(ctx, createPerson,
		arg.FirstName,
		arg.LastName,
		arg.Title,
		arg.Company,
		arg.Age,
		arg.Married,
	)
	return err
}

const findAllPeople = `-- name: FindAllPeople :many
SELECT first_name, last_name, title, company FROM searchable_people limit 30
`

func (q *Queries) FindAllPeople(ctx context.Context) ([]SearchablePerson, error) {
	rows, err := q.db.QueryContext(ctx, findAllPeople)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SearchablePerson
	for rows.Next() {
		var i SearchablePerson
		if err := rows.Scan(
			&i.FirstName,
			&i.LastName,
			&i.Title,
			&i.Company,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
