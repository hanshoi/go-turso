package db

import (
	"context"
	"math/rand"

	"database/sql"
)

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
