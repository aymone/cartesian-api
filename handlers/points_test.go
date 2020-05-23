package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aymone/cartesian-api/handlers"
	"github.com/aymone/cartesian-api/models"
	"github.com/aymone/cartesian-api/router"
	"github.com/aymone/cartesian-api/services"
	"github.com/stretchr/testify/assert"
)

func getURLWithQueryParams(x, y, distance string) string {
	return fmt.Sprintf("%s?x=%s&y=%s&distance=%s", router.URL, x, y, distance)
}

func TestHandlerGetPointsWithSuccess(t *testing.T) {
	data := []*models.Points{
		{X: 5, Y: 2},
		{X: 4, Y: 2},
		{X: 2, Y: 1},
		{X: 3, Y: 1},
		{X: 13, Y: 11},
	}

	s := services.NewPointsService(data)
	h := handlers.New(s)
	r := router.Init(h)

	url := getURLWithQueryParams("3", "2", "10")
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var response []map[string]int
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Len(t, response, 4)
}

func TestHandlerErrorMethodNotAllowed(t *testing.T) {
	h := handlers.New(nil)
	r := router.Init(h)

	url := getURLWithQueryParams("3", "2", "10")
	req, _ := http.NewRequest(http.MethodPost, url, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
}

func TestHandlerGetPointsWithInvalidParams(t *testing.T) {
	h := handlers.New(nil)
	r := router.Init(h)

	// refs: https://github.com/golang/go/wiki/TableDrivenTests
	var tableTests = []struct {
		name string
		url  string
	}{
		{name: "x", url: getURLWithQueryParams("", "2", "")},
		{name: "y", url: getURLWithQueryParams("3", "", "10")},
		{name: "distance", url: getURLWithQueryParams("3", "2", "")},
	}

	for _, tt := range tableTests {
		tt := tt // see refs about table driven tests
		t.Run(tt.name+"_param", func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, tt.url, nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			var response map[string]interface{}
			err := json.Unmarshal([]byte(w.Body.String()), &response)
			assert.Nil(t, err)

			assert.Equal(t, http.StatusBadRequest, w.Code)
			assert.Contains(t, fmt.Sprintf("query parameter '%s' is required", tt.name), response["error"])
			assert.Equal(t, http.StatusBadRequest, int(response["status"].(float64)))
		})
	}
}
