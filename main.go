package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"

	"goh/go-htmx/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	repo "goh/go-htmx/db"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Fprintf(os.Stdout, ".env not foundfile: %s\n", err)
	}

	url := os.Getenv("DB_URL")
	token := os.Getenv("DB_TOKEN")

	if len(token) > 0 {
		url = url + "?authToken=" + token
	}

	if len(url) == 0 {
		fmt.Fprintf(os.Stderr, "No DB url given %s\n", url)
		os.Exit(1)
	}

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s\n", url, err)
		os.Exit(1)
	}
	defer db.Close()

	// create test users for test purposes only
	if repo.CountPeople(db) < 100 {
		repo.CreateTestPeople(db)
	}

	app := echo.New()
	app.Static("/static", "static")
	app.Use(middleware.Recover())
	app.Use(middleware.RemoveTrailingSlash())
	routes.CreateRoutes(app, db)

	log.Fatal(app.Start(":3000"))
}
