package main

import (
	"github.com/nats-io/go-nats"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	nc, _ := nats.Connect(nats.DefaultURL)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}