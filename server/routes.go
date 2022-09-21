package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func StartRoutes() {
	basePath := e.Group("/pc-store")

	basePath.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "PC - Store")
	})
}
