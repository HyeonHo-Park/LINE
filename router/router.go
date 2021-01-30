package router

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

var pingList []pingInfo

func doPing(info pingInfo) {
	// Log File Delete
	logPath := "../pingLog/" + info.Hostname + ".txt"
	os.Remove(logPath)

	// Ping
	p := func(addr string) {
		dst, dur, err := Ping(addr)
		if err != nil {
			f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				log.Fatalf("error opening file: %v", err)
			}
			defer f.Close()
			log.SetOutput(f)
			log.SetFlags(0)
			log.Printf("Ping %s (%s): %s\n", addr, dst, err)
			return
		}
		log.Printf("Ping %s (%s): %s\n", addr, dst, dur)
	}

	for i := 0; i < info.Count; i++ {
		p(info.Hostname)
	}
}

func createPing(c echo.Context) error {
	// Get Values
	hostname := c.FormValue("server")
	count, _ := strconv.Atoi(c.FormValue("count"))
	info := pingInfo{hostname, count}

	// input Ping List
	pingList = append(pingList, info)

	// Make Ping
	go doPing(info)

	return c.JSON(http.StatusOK, &info)
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

	// Delete Ping

	return c.String(http.StatusOK, "Delete Ping "+hostname)
}
