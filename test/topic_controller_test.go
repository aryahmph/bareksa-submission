package test

import (
	"bareksa-aryayunanta/model/domain"
	"bareksa-aryayunanta/repository"
	"context"
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

func TestListTopicSuccess(t *testing.T) {
	db := setupTestDB()
	truncateTopics(db)

	tx, _ := db.Begin()
	topicRepository := repository.NewTopicRepositoryImpl()
	topic1 := topicRepository.Save(context.Background(), tx, domain.Topic{Name: "sports"})
	topic2 := topicRepository.Save(context.Background(), tx, domain.Topic{Name: "technology"})
	tx.Commit()

	router := setupRouter(db)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/topics", nil)
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

	topics := responseBody["data"].([]interface{})

	topicResponse1 := topics[0].(map[string]interface{})
	topicResponse2 := topics[1].(map[string]interface{})

	assert.Equal(t, topic1.Id, uint32(topicResponse1["id"].(float64)))
	assert.Equal(t, topic1.Name, topicResponse1["name"])

	assert.Equal(t, topic2.Id, uint32(topicResponse2["id"].(float64)))
	assert.Equal(t, topic2.Name, topicResponse2["name"])
}
