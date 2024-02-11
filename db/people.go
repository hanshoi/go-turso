package db

import (
	"context"
	"database/sql"
	"log"
	"math/rand"
)

type Person struct {
	Id        uint32
	FirstName string
	LastName  string
	Title     string
	Company   string
	Age       uint8
	Married   bool
}

func GetAllPeople(db *sql.DB) []Person {
	var people []Person
	rows, err := db.Query("SELECT * FROM people")
	if err != nil {
		log.Fatalln(err)
		return people
	}
	defer rows.Close()

	for rows.Next() {
		var person Person
		if err := rows.Scan(&person.Id, &person.FirstName, &person.LastName, &person.Title, &person.Company, &person.Age, &person.Married); err != nil {
			log.Fatalln(err)
			return people
		}
		people = append(people, person)
	}
	return people
}

func SearchPeople(db *sql.DB, keyword string) []Person {
	var people []Person
	rows, err := db.Query("SELECT * FROM searchable_people WHERE searchable_people MATCH ?", keyword)
	if err != nil {
		log.Fatalln(err)
		return people
	}
	defer rows.Close()

	for rows.Next() {
		var person Person
		if err := rows.Scan(&person.Id, &person.FirstName, &person.LastName, &person.Title, &person.Company); err != nil {
			log.Fatalln(err)
			return people
		}
		people = append(people, person)
	}
	return people
}

func CountPeople(db *sql.DB) int {
	count := 0
	err := db.QueryRow("SELECT COUNT(id) FROM people").Scan(&count)
	if err != nil {
		log.Fatalln(err)
	}

	return count
}

func CreateTestPeople(db *sql.DB) {
	query := "INSERT INTO `people` (`first_name`, `last_name`, `title`, `company`, `age`, `married`) VALUES (?, ?, ?, ?, ?, ?)"

	firstnames := []string{"Tim", "John", "James", "Laura", "Chelsea", "Roger", "Charles", "JJ", "Kyle"}
	lastnames := []string{"Smith", "Rogers", "Lemmings", "Tailor", "Butcher", "Priest", "Dalling"}
	titles := []string{"CEO", "CTO", "Peon", "Peasant", "CEO of HTMX", "Priest", "Butcher"}
	companies := []string{"Quteo", "Bad Company"}

	for range 1000 {
		fn := firstnames[rand.Intn(len(firstnames))]
		ln := lastnames[rand.Intn(len(lastnames))]
		title := titles[rand.Intn(len(titles))]
		company := companies[rand.Intn(len(companies))]
		age := rand.Intn(99)
		married := rand.Intn(1)
		_, err := db.ExecContext(context.Background(), query, fn, ln, title, company, age, married)
		if err != nil {
			log.Fatalf("impossible insert person: %s", err)
		}
	}

}
