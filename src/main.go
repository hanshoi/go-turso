package main

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"

	repo "goh/go-htmx/db"
	"goh/go-htmx/routes"

	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	ctx := context.Background()
	settings := LoadSettings()

	url, err := settings.GetURL()
	if err != nil {
		fmt.Fprintf(os.Stderr, "No DB url given. Set ENV var or .env file %d\n", err)
		os.Exit(1)
	}

	conn, err := sql.Open("libsql", url)
	defer conn.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s\n", url, err)
		os.Exit(1)
	}

	db := repo.New(conn)

	// create people for test purposes only
	count, err := db.CountPeople(ctx)
	if err != nil && count < 100 {
		db.CreateTestPeople(ctx)
	}

	app := echo.New()
	app.Static("/static", "static")
	app.Use(middleware.Recover())
	app.Use(middleware.RemoveTrailingSlash())
	routes.CreateRoutes(app, db)

	log.Fatal(app.Start(":3000"))
}
