package test

import (
	"bareksa-aryayunanta/app"
	"bareksa-aryayunanta/controller"
	"bareksa-aryayunanta/helper"
	"bareksa-aryayunanta/middleware"
	"bareksa-aryayunanta/model/domain"
	"bareksa-aryayunanta/repository"
	"bareksa-aryayunanta/service"
	"context"
	"database/sql"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/bareska_aryayunanta_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	tagRepository := repository.NewTagRepositoryImpl()
	tagService := service.NewTagServiceImpl(tagRepository, db, validate)
	tagController := controller.NewTagControllerImpl(tagService)
	router := app.NewRouter(tagController)

	return middleware.NewLogMiddleware(router)
}

func truncateTags(db *sql.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	db.Exec("TRUNCATE tags")
	db.Exec("SET FOREIGN_KEY_CHECKS = 1")
}

func TestCreateTagSuccess(t *testing.T) {
	db := setupTestDB()
	truncateTags(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "random"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/tags", requestBody)
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
	assert.Equal(t, "random", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateTagValidationFailed(t *testing.T) {
	db := setupTestDB()
	truncateTags(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/tags", requestBody)
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

func TestListTagSuccess(t *testing.T) {
	db := setupTestDB()
	truncateTags(db)

	tx, _ := db.Begin()
	tagRepository := repository.NewTagRepositoryImpl()
	tag1 := tagRepository.Save(context.Background(), tx, domain.Tag{Name: "investment"})
	tag2 := tagRepository.Save(context.Background(), tx, domain.Tag{Name: "internship"})
	tx.Commit()

	router := setupRouter(db)
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/tags", nil)
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

	tags := responseBody["data"].([]interface{})

	tagResponse1 := tags[1].(map[string]interface{})
	tagResponse2 := tags[0].(map[string]interface{})

	assert.Equal(t, tag1.ID, uint32(tagResponse1["id"].(float64)))
	assert.Equal(t, tag1.Name, tagResponse1["name"])

	assert.Equal(t, tag2.ID, uint32(tagResponse2["id"].(float64)))
	assert.Equal(t, tag2.Name, tagResponse2["name"])
}
