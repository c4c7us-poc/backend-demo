package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/api/demo/health-check", func(c echo.Context) error {
		fmt.Println("Log health-check in demo")
		log.Info("Log health-check in demo")

		return c.String(http.StatusOK, "it's works")
	})
	e.GET("/api/demo/hello", func(c echo.Context) error {
		fmt.Println("Log api-demo-hello")
		log.Info("Log api-demo-hello")

		return c.String(http.StatusOK, "C4c7us Demo")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
