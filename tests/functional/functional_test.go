package functional

import (
	"bytes"
	"encoding/json"
	"github.com/MasDev-12/mechta.testapi/config"
	"github.com/MasDev-12/mechta.testapi/domain/entities"
	"github.com/MasDev-12/mechta.testapi/infrastructure/db_context"
	"github.com/MasDev-12/mechta.testapi/servers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

var path string = "../../tsconfig.json"

func TestCreateUser(t *testing.T) {
	//Arrange
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Recovered from panic: %v", r)
		}
	}()
	dbSetting, err := config.LoadSettingsDb(path)
	if err != nil {
		panic(err)
	}
	dbContext := db_context.NewDbContext(dbSetting)
	defer dbContext.ClearDatabaseAfterTests()
	serverSetting, err := config.LoadSettingServer(path)
	if err != nil {
		panic(err)
	}
	argon2Setting, err := config.LoadSettingArgon2(path)
	if err != nil {
		panic(err)
	}
	restServer := servers.NewRestServer(serverSetting, dbSetting, argon2Setting)

	router := restServer.Router()

	//Act
	payload := map[string]interface{}{
		"username": "mukanmasud",
		"email":    "example@mail.com",
		"password": "123456",
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/user/create", bytes.NewBuffer(body)) // Изменено на POST
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	//Assert
	assert.Equal(t, http.StatusOK, recorder.Code)

	var response map[string]interface{}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "mukanmasud", response["username"])
	assert.Equal(t, "example@mail.com", response["email"])
	assert.NotNil(t, response["id"])
}

func TestGetUserById(t *testing.T) {
	//Arrange
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Recovered from panic: %v", r)
		}
	}()
	dbSetting, err := config.LoadSettingsDb(path)
	if err != nil {
		panic(err)
	}
	dbContext := db_context.NewDbContext(dbSetting)
	defer dbContext.ClearDatabaseAfterTests()
	serverSetting, err := config.LoadSettingServer(path)
	if err != nil {
		panic(err)
	}
	argon2Setting, err := config.LoadSettingArgon2(path)
	if err != nil {
		panic(err)
	}
	restServer := servers.NewRestServer(serverSetting, dbSetting, argon2Setting)

	router := restServer.Router()
	//Act
	notexistsUserId := "24778205-5d5d-4325-8553-c92955819f33"

	request, _ := http.NewRequest("GET", "/user/"+notexistsUserId, nil)
	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)

	//Assert
	if responseRecorder.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, responseRecorder.Code)
	}
	// Проверить тело ответа (если у вас есть сообщение об ошибке)
	expectedResponse := `{"error":"user not exists"}`
	if responseRecorder.Body.String() != expectedResponse {
		t.Errorf("Expected response body %s, got %s", expectedResponse, responseRecorder.Body.String())
	}
}

func TestCreateUrlSuccess(t *testing.T) {
	//Arrange
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Recovered from panic: %v", r)
		}
	}()
	dbSetting, err := config.LoadSettingsDb(path)
	if err != nil {
		panic(err)
	}
	dbContext := db_context.NewDbContext(dbSetting)
	defer dbContext.ClearDatabaseAfterTests()
	serverSetting, err := config.LoadSettingServer(path)
	if err != nil {
		panic(err)
	}
	argon2Setting, err := config.LoadSettingArgon2(path)
	if err != nil {
		panic(err)
	}
	restServer := servers.NewRestServer(serverSetting, dbSetting, argon2Setting)

	router := restServer.Router()

	user := entities.User{
		Id:           uuid.New(),
		Username:     "test_name",
		Email:        strings.TrimSpace(strings.ToLower("example@mail.com")),
		PasswordHash: "exampleHash",
		IsActive:     true,
		CreatedAt:    time.Now(),
	}
	result, err := restServer.UserCommands.UserService.UserRepository.Add(user)
	if err != nil {
		panic(err)
	}

	//Act
	tempUrl := "example_url"
	payload := map[string]interface{}{
		"original_url": tempUrl,
		"user_id":      result.Id,
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/url/shortener", bytes.NewBuffer(body)) // Изменено на POST
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	//Assert
	var url entities.URL
	urlError := dbContext.Db.Where("origin_url = ?", tempUrl).First(&url).Error
	assert.NoError(t, urlError, "Failed to find the URL in the database")
	assert.NotNil(t, url, "Expected URL to be saved in the database")

	// Проверяем, что short_url был сгенерирован
	assert.NotEmpty(t, url.ShortURL, "Expected short URL to be generated")
}