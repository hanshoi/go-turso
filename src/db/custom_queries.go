package db

import (
	"context"
	"math/rand"

	"database/sql"
)

const searchPeople = `-- name: SearchPeople :many
SELECT first_name, last_name, title, company FROM searchable_people WHERE searchable_people MATCH ? || '*' limit 30
`

func (q *Queries) SearchPeople(ctx context.Context, dollar_1 string) ([]SearchablePerson, error) {
	rows, err := q.db.QueryContext(ctx, searchPeople, dollar_1)
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

func (q *Queries) CreateTestPeople(ctx context.Context) error {
	firstnames := []string{"Tim", "John", "James", "Laura", "Chelsea", "Roger", "Charles", "JJ", "Kyle"}
	lastnames := []string{"Smith", "Rogers", "Lemmings", "Tailor", "Butcher", "Priest", "Dalling"}
	titles := []string{"CEO", "CTO", "Peon", "Peasant", "CEO of HTMX", "Priest", "Butcher"}
	companies := []string{"Quteo", "Bad Company"}

	for i := 1; i <= 10000; i++ {
		params := CreatePersonParams{
			firstnames[rand.Intn(len(firstnames))],
			lastnames[rand.Intn(len(lastnames))],
			sql.NullString{String: titles[rand.Intn(len(titles))]},
			sql.NullString{String: companies[rand.Intn(len(companies))]},
			sql.NullInt64{Int64: int64(rand.Intn(99))},
			rand.Intn(2) == 1,
		}
		err := q.CreatePerson(ctx, params)
		if err != nil {
			return err
		}
	}
	return nil
}
