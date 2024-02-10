package utils

import (
	"context"
	"goh/go-htmx/templates"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(ctx echo.Context, status int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(status)

	err := t.Render(context.Background(), ctx.Response().Writer)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "failed to render response template")
	}

	return nil
}

func RenderPage(ctx echo.Context, status int, page templates.Page, t templ.Component) error {
	cc := &HtmxContext{Context: ctx}
	content := templates.Content(t, templates.NavBar(page))
	if cc.IsHtmx() {
		return Render(ctx, status, content)
	}
	return Render(ctx, status, templates.Base(content))
}
