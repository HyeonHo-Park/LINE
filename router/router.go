package router

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	. "github.com/HyeonHo-Park/LINE/handler"
	. "github.com/HyeonHo-Park/LINE/model"
	. "github.com/HyeonHo-Park/LINE/utils"
	"github.com/hpcloud/tail"
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
	index := GetIndexByHostname(pingList, hostname)
	fileLength := pingList[index].Count

	if index == -1 {
		return c.String(http.StatusOK, "이미 수행이 끝난 작업 입니다.")
	} else {
		if wait != "true" {
			// return current job result
			return c.String(http.StatusOK, ReadPingLog(hostname))
		} else {
			// return job result until done
			logPath := "/tmp/pingLog/" + hostname + ".txt"
			// logPath := "/Users/hyeonho/Desktop/LINE/pingLog/" + hostname + ".txt"

			// set response
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c.Response().WriteHeader(http.StatusOK)
			enc := json.NewEncoder(c.Response())

			// open file
			t, _ := tail.TailFile(logPath, tail.Config{Follow: true})

			// read file
			for line := range t.Lines {
				// encode text
				if err := enc.Encode(line); err != nil {
					return err
				}
				// send Message
				c.Response().Flush()
				time.Sleep(30 * time.Millisecond)

				// check file
				icmpSeq := strings.Split(line.Text, " ")
				seq := strings.Split(icmpSeq[3], "=")
				if seq[1] == strconv.Itoa(fileLength) {
					break
				}
			}
			return nil
		}
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

func Health(c echo.Context) error {
	return c.String(http.StatusOK, "UP")
}
