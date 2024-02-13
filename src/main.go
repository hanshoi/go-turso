package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"

	"goh/go-htmx/db"
	"goh/go-htmx/routes"
	"goh/go-htmx/turso"
	"goh/go-htmx/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	settings := utils.LoadSettings()

	url, err := settings.DB.GetURL()
	if err != nil {
		fmt.Fprintf(os.Stderr, "No DB url given. Set ENV var or .env file %d\n", err)
		os.Exit(1)
	}

	api := turso.CreateApi(settings)
	api.GetLocations()

	connection, err := db.CreateConnection(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s\n", url, err)
		os.Exit(1)
	}

	// create test users for test purposes only
	if db.CountPeople(connection) < 100 {
		db.CreateTestPeople(connection)
	}

	app := echo.New()
	app.Static("/static", "static")
	app.Use(middleware.Recover())
	app.Use(middleware.RemoveTrailingSlash())
	routes.CreateRoutes(app, connection)

	log.Fatal(app.Start(":3000"))
}
