package main_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	. "github.com/HyeonHo-Park/LINE/model"
)

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

const (
	hostname1 = "google.com"
	hostname2 = "naver.com"
	count1    = 100
	count2    = 200
)

func TestCreatePing(t *testing.T) {
	// Set UP
	e := echo.New()
	formData := PingInfo{hostname1, count1}

	f := make(url.Values)
	f.Set("server", formData.Hostname)
	f.Set("count", strconv.Itoa(formData.Count))

	// Act
	req := httptest.NewRequest(http.MethodPost, "/ping", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	fmt.Println(c)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, formData, c.Request().GetBody)
}

/*
	# Description : 수행중인 Ping 작업 조회

	# Request
		GET /
		[
			{"hostname": "google.com", "count": 100},
			{"hostname": "naver.com", "count": 200}
		]
*/
func TestGetPingList(t *testing.T) {
	// Set UP
	e := echo.New()
	formData1 := PingInfo{hostname1, count1}
	formData2 := PingInfo{hostname2, count2}

	var pingList []PingInfo
	pingList = append(pingList, formData1)
	pingList = append(pingList, formData2)

	// Act
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	fmt.Println(c.Response())
	fmt.Println(rec.Body.String())
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, pingList, rec.Body.String())
}

/*
	# Description : Ping 작업 삭제

	# Request
		DELETE /google.com

	# Response
		Content-Type: text/plain
*/
func TestDeletePing(t *testing.T) {
	// Set UP
	e := echo.New()
	formData1 := PingInfo{hostname1, count1}
	formData2 := PingInfo{hostname2, count2}

	var pingList []PingInfo
	pingList = append(pingList, formData1)
	pingList = append(pingList, formData2)

	// Act
	req := httptest.NewRequest(http.MethodDelete, "/", strings.NewReader(hostname1))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	fmt.Println(c)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, pingList, rec.Body.String())
}
