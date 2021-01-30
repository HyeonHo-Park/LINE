package router

import (
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

	if CheckByHostname(pingList, hostname) {
		return c.JSON(http.StatusOK, "이미 체크 중인 Server입니다.")
	} else {
		// input Ping List
		pingList = append(pingList, info)

		// Make Ping
		go DoPing(&pingList, info)
		return c.JSON(http.StatusOK, &info)
	}
}

func GetPing(c echo.Context) error {
	// get hostname in param
	hostname := c.Param("hostname")
	wait := c.QueryParam("wait")

	if wait != "true" {
		// return current job result
		return c.String(http.StatusOK, ReadPingLog(hostname))
	} else {
		// return job result until done

		// get HTTP Connection

		// send message
		return c.String(http.StatusOK, hostname)
	}
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
