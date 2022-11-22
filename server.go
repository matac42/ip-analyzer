package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/matac42/ip-analyzer/analyzer"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/analyze", getResult)
	e.Logger.Fatal(e.Start(":1323"))
}

func getResult(c echo.Context) error {
	return c.String(http.StatusOK, analyzer.Analyzer())
}
