package main

import (
	. "github.com/HyeonHo-Park/LINE/router"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	e.POST("/ping", CreatePing)

	e.GET("/:hostname", GetPing)

	e.GET("/", GetPingList)

	e.DELETE("/:hostname", DeletePing)

	e.Logger.Fatal(e.Start(":1323"))
}
