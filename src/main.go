package main

import (
	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger TEST API
// @version 1.0

// @contact.name Hyeonho Park
// @contact.url https://github.com/HyeonHo-Park/LINE
// @contact.email hyunho129@naver.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/ping", createPing)

	e.GET("/:hostname", getPing)

	e.GET("/", getPingList)

	e.DELETE("/:hostname", deletePing)

	e.Logger.Fatal(e.Start(":1323"))
}
