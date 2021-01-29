package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/HyeonHo-Park/LINE-TEST/src/server/docs"
)

// @title Swagger TEST API
// @version 1.0

// @contact.name Hyeonho Park
// @contact.url https://github.com/HyeonHo-Park/LINE-TEST
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

/*
	# Description : Ping 작업 생성

	# Request
		POST /ping
		Content-Type: multipart/form-data
		server=google.com&count=100

	# Response
		Content-Type: application/json
		{"hostname": "google.com", "count": 100}
*/
func createPing(c echo.Context) error {
	// def result
	var result struct {
		Hostname string `json:"hostname"`
		Count    int    `json:"count"`
	}

	// Get Values
	result.Hostname = c.FormValue("server")
	result.Count, _ = strconv.Atoi(c.FormValue("count"))

	// Make Ping
	fmt.Println("Create Job")

	return c.JSON(http.StatusOK, &result)
	// c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	// c.Response().WriteHeader(http.StatusOK)
	// return json.NewEncoder(c.Response()).Encode(result)
}

/*
	# Description
		특정 Ping output return

	# Request
		GET /google.com

	# Request
		GET /google.com?wait=true

	# Response
		Content-Type: text/plain
		PING google.com (172.217.26.110): 56 data bytes
		64 bytes from 172.217.26.110: icmp_seq=0 ttl=51
		64 bytes from 172.217.26.110: icmp_seq=1 ttl=51
		64 bytes from 172.217.26.110: icmp_seq=2 ttl=51

		만약 wait=true인 경우, 종료될때까지 HTTP Connection이 끊이지 않은채로 전달되어야 한다.
*/
func getPing(c echo.Context) error {
	// get hostname in param
	hostname := c.Param("hostname")
	wait := c.QueryParam("wait")

	// check parameter
	fmt.Println(hostname, wait)

	// get HTTP Connection

	// send message

	return c.String(http.StatusOK, hostname)
}

/*
	# Description : 수행중인 Ping 작업 조회

	# Request
		GET /
*/
func getPingList(c echo.Context) error {
	return c.String(http.StatusOK, "Ping List")
}

/*
	# Description : Ping 작업 삭제

	# Request
		DELETE /google.com

	# Response
		Content-Type: text/plain
*/
func deletePing(c echo.Context) error {
	// get hostname in param
	hostname := c.Param("hostname")

	// check parameter
	fmt.Println(hostname)

	return c.String(http.StatusOK, "Delete Ping")
}
