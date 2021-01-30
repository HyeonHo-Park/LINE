package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	. "github.com/HyeonHo-Park/LINE/model"
	. "github.com/HyeonHo-Park/LINE/router"
)

const (
	hostname1 = "google.com"
	hostname2 = "naver.com"
	count1    = 100
	count2    = 200
)

func TestCreatePing(t *testing.T) {
	// Set UP
	e := echo.New()
	e.POST("/ping", CreatePing)

	case1 := PingInfo{hostname1, count1}
	f := make(url.Values)
	f.Set("server", hostname1)
	f.Set("count", strconv.Itoa(count1))

	// Act
	req := httptest.NewRequest(http.MethodPost, "/ping", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	// Assertions
	if assert.Equal(t, http.StatusOK, rec.Code) {
		var resBody PingInfo
		json.Unmarshal([]byte(rec.Body.String()), &resBody)
		assert.Equal(t, case1, resBody)
	}
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
func TestGetPing(t *testing.T) {

}

func TestGetPingList(t *testing.T) {
	// Set UP
	e := echo.New()
	e.POST("/ping", CreatePing)
	e.GET("/", GetPingList)

	formData1 := PingInfo{hostname1, count1}
	formData2 := PingInfo{hostname2, count2}

	var pingList []PingInfo
	pingList = append(pingList, formData1)
	pingList = append(pingList, formData2)

	// add to server formData1
	f := make(url.Values)
	f.Set("server", hostname1)
	f.Set("count", strconv.Itoa(count1))

	req := httptest.NewRequest(http.MethodPost, "/ping", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	// add to server formData2
	f2 := make(url.Values)
	f2.Set("server", hostname2)
	f2.Set("count", strconv.Itoa(count2))

	req = httptest.NewRequest(http.MethodPost, "/ping", strings.NewReader(f2.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	// Act
	req = httptest.NewRequest(http.MethodGet, "/", nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	// Assertions
	if assert.Equal(t, http.StatusOK, rec.Code) {
		var resBody []PingInfo
		json.Unmarshal([]byte(rec.Body.String()), &resBody)
		assert.Equal(t, pingList, resBody)
	}
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
	e.POST("/ping", CreatePing)
	e.GET("/", GetPingList)
	e.DELETE("/:hostname", DeletePing)

	formData2 := PingInfo{hostname2, count2}

	var pingList []PingInfo
	pingList = append(pingList, formData2)

	// add to server formData1
	f := make(url.Values)
	f.Set("server", hostname1)
	f.Set("count", strconv.Itoa(count1))

	req := httptest.NewRequest(http.MethodPost, "/ping", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	// add to server formData2
	f2 := make(url.Values)
	f2.Set("server", hostname2)
	f2.Set("count", strconv.Itoa(count2))

	req = httptest.NewRequest(http.MethodPost, "/ping", strings.NewReader(f2.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	// Act
	req = httptest.NewRequest(http.MethodDelete, "/"+hostname1, nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	// Get List
	req = httptest.NewRequest(http.MethodGet, "/", nil)
	rec = httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	// Assertions
	if assert.Equal(t, http.StatusOK, rec.Code) {
		var resBody []PingInfo
		json.Unmarshal([]byte(rec.Body.String()), &resBody)
		assert.Equal(t, pingList, resBody)
	}
}
