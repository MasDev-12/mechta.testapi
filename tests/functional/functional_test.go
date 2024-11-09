package functional

import (
	"bytes"
	"encoding/json"
	"github.com/MasDev-12/mechta.testapi/application/CQRS/responses"
	"github.com/MasDev-12/mechta.testapi/config"
	"github.com/MasDev-12/mechta.testapi/domain/entities"
	"github.com/MasDev-12/mechta.testapi/infrastructure/db_context"
	"github.com/MasDev-12/mechta.testapi/servers"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

var path string = "../../test_tsconfig.json"

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
	mockServer := servers.NewMockRestServer(serverSetting, dbSetting, argon2Setting)
	if err := mockServer.StartMockServer(); err != nil {
		log.Fatalf("Failed to start REST server: %v", err)
	}

	router := mockServer.Router()

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
	mockServer := servers.NewMockRestServer(serverSetting, dbSetting, argon2Setting)
	if err := mockServer.StartMockServer(); err != nil {
		log.Fatalf("Failed to start REST server: %v", err)
	}

	router := mockServer.Router()
	//Act
	notexistsUserId := "24778205-5d5d-4325-8553-c92955819f33"

	request, _ := http.NewRequest("GET", "/user/"+notexistsUserId, nil)
	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)

	//Assert
	if responseRecorder.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, responseRecorder.Code)
	}

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
	mockServer := servers.NewMockRestServer(serverSetting, dbSetting, argon2Setting)
	if err := mockServer.StartMockServer(); err != nil {
		log.Fatalf("Failed to start REST server: %v", err)
	}

	router := mockServer.Router()

	user := entities.User{
		Id:           uuid.New(),
		Username:     "test_name",
		Email:        strings.TrimSpace(strings.ToLower("example@mail.com")),
		PasswordHash: "exampleHash",
		IsActive:     true,
		CreatedAt:    time.Now(),
	}
	result, err := mockServer.UserCommands.UserService.UserRepository.Add(user)
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

func TestGetUserUrlSuccess(t *testing.T) {
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
	mockServer := servers.NewMockRestServer(serverSetting, dbSetting, argon2Setting)
	if err := mockServer.StartMockServer(); err != nil {
		log.Fatalf("Failed to start REST server: %v", err)
	}

	router := mockServer.Router()

	user := entities.User{
		Id:           uuid.New(),
		Username:     "test_name",
		Email:        strings.TrimSpace(strings.ToLower("example@mail.com")),
		PasswordHash: "exampleHash",
		IsActive:     true,
		CreatedAt:    time.Now(),
	}
	result, err := mockServer.UserCommands.UserService.UserRepository.Add(user)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 2; i++ {
		url := entities.URL{
			Id:             uuid.New(),
			OriginalURL:    "example_url_" + strconv.Itoa(i),
			ShortURL:       "exm_url" + strconv.Itoa(i),
			UserId:         result.Id,
			CreatedAt:      time.Now(),
			IsActive:       true,
			ExpiresAt:      time.Now().Add(24 * time.Hour),
			ClickCount:     i,
			LastAccessedAt: nil,
		}
		_, UrlError := mockServer.URLCommands.URLService.URLRepository.Add(url)
		if UrlError != nil {
			panic(err)
		}
	}

	//Act
	request, _ := http.NewRequest("GET", "/url/shortener/"+result.Id.String(), nil)
	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)

	//Assert
	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Expected HTTP status code to be 200 OK")

	var response responses.GetUserUrlsResponse
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &response)
	assert.NoError(t, err, "Failed to parse response body")

	assert.Len(t, response.Urls, 2, "Expected 2 URLs to be returned")

	for i, url := range response.Urls {
		expectedOriginalURL := "example_url_" + strconv.Itoa(i)
		expectedShortURL := "exm_url" + strconv.Itoa(i)

		assert.Equal(t, expectedOriginalURL, url.OriginalURL, "Unexpected Original URL at index %d", i)
		assert.Equal(t, expectedShortURL, url.ShortURL, "Unexpected Short URL at index %d", i)
	}
}

func TestGetUrlByLinkSuccess(t *testing.T) {
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
	mockServer := servers.NewMockRestServer(serverSetting, dbSetting, argon2Setting)
	if err := mockServer.StartMockServer(); err != nil {
		log.Fatalf("Failed to start REST server: %v", err)
	}

	router := mockServer.Router()

	user := entities.User{
		Id:           uuid.New(),
		Username:     "test_name",
		Email:        strings.TrimSpace(strings.ToLower("example@mail.com")),
		PasswordHash: "exampleHash",
		IsActive:     true,
		CreatedAt:    time.Now(),
	}
	result, err := mockServer.UserCommands.UserService.UserRepository.Add(user)
	if err != nil {
		panic(err)
	}
	urlEntity := entities.URL{
		Id:             uuid.New(),
		OriginalURL:    "example_url_",
		ShortURL:       "exm_url",
		UserId:         result.Id,
		CreatedAt:      time.Now(),
		IsActive:       true,
		ExpiresAt:      time.Now().Add(24 * time.Hour),
		ClickCount:     0,
		LastAccessedAt: nil, // Примерный срок действия
	}
	url, UrlError := mockServer.URLCommands.URLService.URLRepository.Add(urlEntity)
	if UrlError != nil {
		panic(err)
	}
	//Act
	request, _ := http.NewRequest("GET", "/url/"+url.ShortURL, nil)
	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)

	//Assert
	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Expected HTTP status code to be 200 OK")

	var responseBody responses.GetUrlByShortNameResponse
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &responseBody)
	assert.NoError(t, err, "Expected no error during JSON unmarshalling")
	assert.NotNil(t, responseBody.Url, "Expected Url field to be non-nil")
	assert.Nil(t, responseBody.Error, "Expected Error field to be nil")

	assert.Equal(t, urlEntity.Id, responseBody.Url.Id, "Expected returned URL ID to match")
	assert.Equal(t, urlEntity.OriginalURL, responseBody.Url.OriginalURL, "Expected OriginalURL to match")
	assert.Equal(t, urlEntity.ShortURL, responseBody.Url.ShortURL, "Expected ShortURL to match")
	assert.Equal(t, urlEntity.UserId, responseBody.Url.UserId, "Expected UserId to match")
	assert.Equal(t, urlEntity.IsActive, responseBody.Url.IsActive, "Expected IsActive to match")
	assert.Equal(t, urlEntity.ClickCount+1, responseBody.Url.ClickCount, "Expected ClickCount to match")
}

func TestDeleteUrlByLinkSuccess(t *testing.T) {
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
	mockServer := servers.NewMockRestServer(serverSetting, dbSetting, argon2Setting)
	if err := mockServer.StartMockServer(); err != nil {
		log.Fatalf("Failed to start REST server: %v", err)
	}

	router := mockServer.Router()

	user := entities.User{
		Id:           uuid.New(),
		Username:     "test_name",
		Email:        strings.TrimSpace(strings.ToLower("example@mail.com")),
		PasswordHash: "exampleHash",
		IsActive:     true,
		CreatedAt:    time.Now(),
	}
	result, err := mockServer.UserCommands.UserService.UserRepository.Add(user)
	if err != nil {
		panic(err)
	}
	urlEntity := entities.URL{
		Id:             uuid.New(),
		OriginalURL:    "example_url_",
		ShortURL:       "exm_url",
		UserId:         result.Id,
		CreatedAt:      time.Now(),
		IsActive:       true,
		ExpiresAt:      time.Now().Add(24 * time.Hour),
		ClickCount:     0,
		LastAccessedAt: nil, // Примерный срок действия
	}
	url, UrlError := mockServer.URLCommands.URLService.URLRepository.Add(urlEntity)
	if UrlError != nil {
		panic(err)
	}

	//Act
	request, _ := http.NewRequest("DELETE", "/url/"+url.ShortURL, nil)
	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)

	//Assert
	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Expected HTTP status code to be 200 OK")

	deletedUrl, shortUrlErr := mockServer.URLQueries.URLService.URLRepository.GetUrlByShortName(url.ShortURL)
	assert.NoError(t, shortUrlErr, "Expected no error while fetching URL by short name")
	assert.Nil(t, deletedUrl, "Expected the deleted URL to be nil in the database")
}

func TestGetLinkStatSuccess(t *testing.T) {
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
	mockServer := servers.NewMockRestServer(serverSetting, dbSetting, argon2Setting)
	if err := mockServer.StartMockServer(); err != nil {
		log.Fatalf("Failed to start REST server: %v", err)
	}

	router := mockServer.Router()

	user := entities.User{
		Id:           uuid.New(),
		Username:     "test_name",
		Email:        strings.TrimSpace(strings.ToLower("example@mail.com")),
		PasswordHash: "exampleHash",
		IsActive:     true,
		CreatedAt:    time.Now(),
	}
	result, err := mockServer.UserCommands.UserService.UserRepository.Add(user)
	if err != nil {
		panic(err)
	}
	urlEntity := entities.URL{
		Id:             uuid.New(),
		OriginalURL:    "example_url_",
		ShortURL:       "exm_url",
		UserId:         result.Id,
		CreatedAt:      time.Now(),
		IsActive:       true,
		ExpiresAt:      time.Now().Add(24 * time.Hour),
		ClickCount:     0,
		LastAccessedAt: nil, // Примерный срок действия
	}
	url, UrlError := mockServer.URLCommands.URLService.URLRepository.Add(urlEntity)
	if UrlError != nil {
		panic(err)
	}

	//Act
	request, _ := http.NewRequest("GET", "/url/stats/"+url.ShortURL, nil)
	responseRecorder := httptest.NewRecorder()
	router.ServeHTTP(responseRecorder, request)

	//Assert
	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Expected HTTP status code to be 200 OK")
	var responseBody responses.GetUrlStatByShortNameResponse
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &responseBody)
	assert.NoError(t, err, "Expected no error during JSON unmarshalling")
	assert.NotNil(t, responseBody, "Expected Url field to be non-nil")
	assert.Nil(t, responseBody.Error, "Expected Error field to be nil")

	assert.Equal(t, urlEntity.OriginalURL, *responseBody.OriginalURL, "Expected OriginalURL to match")
	assert.Equal(t, urlEntity.IsActive, *responseBody.IsActive, "Expected IsActive to match")
	assert.Equal(t, urlEntity.ClickCount+1, *responseBody.ClickCount, "Expected ClickCount to match")
}
