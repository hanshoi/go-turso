package db

import (
	"database/sql"
	"log"
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

func GetPeople(db *sql.DB) []Person {
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
