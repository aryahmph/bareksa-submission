package test

import (
	"database/sql"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func truncateTopics(db *sql.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("TRUNCATE topics")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func TestCreateTopicSuccess(t *testing.T) {
	db := setupTestDB()
	truncateTopics(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "technology"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/topics", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "BAREKSA_INTERNSHIP")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "technology", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateTopicValidationFailed(t *testing.T) {
	db := setupTestDB()
	truncateTopics(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/topics", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "BAREKSA_INTERNSHIP")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}
