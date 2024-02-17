-- name: FindAllPeople :many
SELECT * FROM searchable_people limit 30;

-- name: SearchPeople :many
SELECT * FROM searchable_people WHERE searchable_people MATCH ? || '*' limit 30;

-- name: CountPeople :one
SELECT COUNT(id) FROM people;

-- name: CreatePerson :exec
INSERT INTO people (first_name, last_name, title, company, age, married) VALUES (?, ?, ?, ?, ?, ?);
