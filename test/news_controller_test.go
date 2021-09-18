package test

import (
	"bytes"
	"database/sql"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func truncateNews(db *sql.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("TRUNCATE news")
	db.Exec("TRUNCATE news_tags")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

//go:embed img/bareksa.jpg
var uploadFileTest []byte

func TestCreateNewsSuccess(t *testing.T) {
	db := setupTestDB()
	truncateTopics(db)
	truncateTags(db)
	truncateNews(db)
	router := setupRouter(db)

	// Add Topics
	requestBodyTopic := strings.NewReader(`{"name" : "business"}`)
	requestTopic := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/topics", requestBodyTopic)
	requestTopic.Header.Add("Content-Type", "application/json")
	requestTopic.Header.Add("X-API-Key", "BAREKSA_INTERNSHIP")

	recorderTopic := httptest.NewRecorder()
	router.ServeHTTP(recorderTopic, requestTopic)

	requestBody := new(bytes.Buffer)
	writer := multipart.NewWriter(requestBody)
	writer.WriteField("title", "Bareksa membuka kesempatan magang")
	writer.WriteField("topic_name", "business")
	writer.WriteField("short_desc", "lorem ipsum")
	writer.WriteField("content", "lorem ipsum sit dolor amet hehe")
	writer.WriteField("writer", "Arya Yunanta")
	writer.WriteField("tags", "internship,backend")

	file, _ := writer.CreateFormFile("image", "EXAMPLE.jpg")
	file.Write(uploadFileTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/news", requestBody)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	request.Header.Add("X-API-Key", "BAREKSA_INTERNSHIP")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	fmt.Println(response)
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	assert.Equal(t, "Bareksa membuka kesempatan magang", responseBody["data"].(map[string]interface{})["title"])
	assert.Equal(t, "business", responseBody["data"].(map[string]interface{})["topic"])
	assert.Equal(t, "lorem ipsum", responseBody["data"].(map[string]interface{})["description"])
	assert.Equal(t, "lorem ipsum sit dolor amet hehe", responseBody["data"].(map[string]interface{})["content"])
	assert.Equal(t, "Arya Yunanta", responseBody["data"].(map[string]interface{})["writer"])
	assert.Equal(t, "internship,backend", responseBody["data"].(map[string]interface{})["tags"])
}
