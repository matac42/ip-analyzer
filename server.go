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
	e.GET("/analyze", executeAnalyze)
	e.GET("/result", getResult)
	e.GET("/progress", getProgress)
	e.Logger.Fatal(e.Start(":1323"))
}

func executeAnalyze(c echo.Context) error {
	return c.String(http.StatusOK, analyzer.Analyze())
}

func getResult(c echo.Context) error {
	return c.String(http.StatusOK, analyzer.ArpResult())
}

func getProgress(c echo.Context) error {
	return c.String(http.StatusOK, analyzer.GetProgress())
}
