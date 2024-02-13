package utils

import (
	"github.com/labstack/echo/v4"
)

type HtmxContext struct {
	echo.Context
}

func (h *HtmxContext) IsHtmx() bool {
	return h.Request().Header.Get("Hx-Request") != ""
}
