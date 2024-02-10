package routes

import (
	"net/http"

	repo "goh/go-htmx/db"
	"goh/go-htmx/templates"
	"goh/go-htmx/utils"

	"database/sql"
	"github.com/labstack/echo/v4"
)

type MyHandlerFunction func(echo.Context, *sql.DB) error

func CreateRoutes(app *echo.Echo, db *sql.DB) {
	app.GET("/", wrap(people, db))
	app.GET("/detail/", wrap(detail, db))
	app.GET("/about/", about)
}

func wrap(handler MyHandlerFunction, db *sql.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return handler(ctx, db)
	}
}

func people(ctx echo.Context, db *sql.DB) error {
	// cc := &utils.HtmxContext{Context: ctx}
	// name := "Stacy"
	// if cc.IsHtmx() {
	// 	name = "Randy"
	// }

	people := repo.GetPeople(db)
	if len(people) == 0 {
		return utils.RenderPage(ctx, http.StatusOK, templates.PeoplePageType, templates.EmptyListPage())
	}
	return utils.RenderPage(ctx, http.StatusOK, templates.PeoplePageType, templates.PeoplePage(people))
}

func detail(ctx echo.Context, db *sql.DB) error {
	return utils.RenderPage(ctx, http.StatusOK, templates.PeoplePageType, templates.DetailPage())
}

func about(ctx echo.Context) error {
	return utils.RenderPage(ctx, http.StatusOK, templates.AboutPageType, templates.AboutPage())
}
