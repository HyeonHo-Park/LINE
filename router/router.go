package router

import (
	"fmt"
	"net/http"
	"strconv"

	. "github.com/HyeonHo-Park/LINE/handler"
	. "github.com/HyeonHo-Park/LINE/model"
	. "github.com/HyeonHo-Park/LINE/utils"
	"github.com/labstack/echo/v4"
)

var pingList []PingInfo

func CreatePing(c echo.Context) error {
	// Get Values
	hostname := c.FormValue("server")
	count, _ := strconv.Atoi(c.FormValue("count"))
	info := PingInfo{hostname, count}

	// input Ping List
	pingList = append(pingList, info)

	// Make Ping
	go doPing(info)

	return c.JSON(http.StatusOK, &info)
}

func GetPing(c echo.Context) error {
	// get hostname in param
	hostname := c.Param("hostname")
	wait := c.QueryParam("wait")

	// check parameter
	fmt.Println(hostname, wait)

	// return current job result

	// return job result until done
	// get HTTP Connection

	// send message

	return c.String(http.StatusOK, hostname)
}

func GetPingList(c echo.Context) error {
	return c.JSON(http.StatusOK, pingList)
}

func DeletePing(c echo.Context) error {
	// get hostname in param
	hostname := c.Param("hostname")

	// delete hostname in pingList
	pingList = RemoveByHostname(pingList, hostname)

	// Delete Ping

	return c.String(http.StatusOK, "Delete Ping "+hostname)
}
