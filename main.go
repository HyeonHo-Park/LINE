package main

import (
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	e.POST("/ping", createPing)

	e.GET("/:hostname", getPing)

	e.GET("/", getPingList)

	e.DELETE("/:hostname", deletePing)

	e.Logger.Fatal(e.Start(":1323"))
}
