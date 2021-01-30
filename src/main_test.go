package test

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

type pingInfo struct {
	Hostname string `json:"hostname"`
	Count    int    `json:"count"`
}

func TestCreatePing(t *testing.T) {
	// Set UP
	e := echo.New()
	formData := pingInfo{"google.com", 100}

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
	assert.Equal(t, formData, rec.Body.String())
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
	formData1 := pingInfo{"google.com", 100}
	formData2 := pingInfo{"naver.com", 200}

	var pingList []pingInfo
	pingList = append(pingList, formData1)
	pingList = append(pingList, formData2)

	// Act
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	fmt.Println(c)
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

	// Act

	// Assertions
}
