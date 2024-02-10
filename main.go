package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// "github.com/golang-migrate/migrate/v4"
	// "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/tursodatabase/libsql-client-go/libsql"

	"goh/go-htmx/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	url := "http://127.0.0.1:8080"

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}
	defer db.Close()

	// driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	// m, err := migrate.NewWithDatabaseInstance(
	// 	"file:///db/migrations",
	// 	"sqlite", driver)
	// m.Up()

	app := echo.New()
	app.Static("/static", "static")
	app.Use(middleware.Recover())
	app.Use(middleware.RemoveTrailingSlash())
	routes.CreateRoutes(app, db)

	log.Fatal(app.Start(":3000"))
}
