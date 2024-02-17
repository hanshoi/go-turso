package routes

import (
	"net/http"

	repo "goh/go-htmx/db"
	"goh/go-htmx/templates"
	"goh/go-htmx/utils"

	"github.com/labstack/echo/v4"
	"log"
)

type MyHandlerFunction func(echo.Context, *repo.Queries) error

func CreateRoutes(app *echo.Echo, db *repo.Queries) {
	app.GET("/", wrap(people, db))
	app.POST("/search/", wrap(search, db))
	app.GET("/detail/:id", wrap(detail, db))
	app.GET("/about/", about)
}

func wrap(handler MyHandlerFunction, db *repo.Queries) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return handler(ctx, db)
	}
}

func search(ctx echo.Context, db *repo.Queries) error {
	keyword := ctx.FormValue("search")
	var people []repo.SearchablePerson
	var err error
	if len(keyword) > 0 {
		people, err = db.SearchPeople(ctx.Request().Context(), keyword)

	} else {
		people, err = db.FindAllPeople(ctx.Request().Context())
	}

	if err != nil {
		log.Fatalln(err)
		return err
	}

	return utils.Render(ctx, http.StatusOK, templates.ListPeople(people))
}

func people(ctx echo.Context, db *repo.Queries) error {
	people, err := db.FindAllPeople(ctx.Request().Context())

	if err != nil || len(people) == 0 {
		return utils.RenderPage(ctx, http.StatusOK, templates.PeoplePageType, templates.EmptyListPage())
	}
	return utils.RenderPage(ctx, http.StatusOK, templates.PeoplePageType, templates.PeoplePage(people))
}

func detail(ctx echo.Context, db *repo.Queries) error {
	return utils.RenderPage(ctx, http.StatusOK, templates.PeoplePageType, templates.DetailPage())
}

func about(ctx echo.Context) error {
	return utils.RenderPage(ctx, http.StatusOK, templates.AboutPageType, templates.AboutPage())
}
