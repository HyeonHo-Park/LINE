package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type pingInfo struct {
	Hostname string `json:"hostname"`
	Count    int    `json:"count"`
}

var pingList []pingInfo

func createPing(c echo.Context) error {
	// Get Values
	hostname := c.FormValue("server")
	count, _ := strconv.Atoi(c.FormValue("count"))
	result := pingInfo{hostname, count}

	// input Ping List
	pingList = append(pingList, result)

	// Make Job
	p := func(addr string) {
		dst, dur, err := Ping(addr)
		if err != nil {
			log.Printf("Ping %s (%s): %s\n", addr, dst, err)
			return
		}
		log.Printf("Ping %s (%s): %s\n", addr, dst, dur)
	}
	for i := 0; i < result.Count; i++ {
		p(result.Hostname)
	}

	return c.JSON(http.StatusOK, &result)
}

func getPing(c echo.Context) error {
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

func getPingList(c echo.Context) error {
	return c.JSON(http.StatusOK, pingList)
}

func deletePing(c echo.Context) error {
	// get hostname in param
	hostname := c.Param("hostname")

	// delete hostname in pingList
	RemoveByHostname(hostname)

	// delete job

	return c.String(http.StatusOK, "Delete Ping "+hostname)
}
