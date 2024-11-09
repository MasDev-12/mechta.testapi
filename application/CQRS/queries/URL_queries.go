package queries

import (
	"github.com/MasDev-12/mechta.testapi/application/CQRS/requests"
	"github.com/MasDev-12/mechta.testapi/application/CQRS/responses"
	"github.com/MasDev-12/mechta.testapi/application/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
)

type URLQueries struct {
	URLService *services.URLService
}

func NewURLQueries(urlService *services.URLService) *URLQueries {
	return &URLQueries{
		URLService: urlService,
	}
}

// GetUserUrls godoc
// @Summary Get all URLs created by a user
// @Description Get all shortened URLs by providing a user ID in the request URL
// @Tags URLs
// @Accept  json
// @Produce  json
// @Param userId path string true "User ID"
// @Success 200 {object} responses.GetUserUrlsResponse "Successfully retrieved user URLs"
// @Failure 400 {object} string "Invalid user ID"
// @Failure 404 {object} string "No URLs found for the user"
// @Failure 408 {object} string "Request timed out"
// @Failure 500 {object} string "Internal Server Error"
// @Router /url/shortener/{userId} [get]
func (query *URLQueries) GetUserUrls(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid user Id"})
		return
	}
	responseChan := make(chan responses.GetUserUrlsResponse)
	timeout := time.After(10 * time.Second)

	go func() {
		responseChan <- query.URLService.GetUserUrls(requests.GetUserUrlsRequest{
			UserId: userId.(uuid.UUID),
		})
	}()

	select {
	case response := <-responseChan:
		if response.Error != nil {
			if strings.Contains(response.Error.Error(), "url not found") {
				c.JSON(http.StatusNotFound, gin.H{"Error": response.Error.Error()})
				return // 404 если url не найден
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"Error": response.Error.Error()})
				return // Другие ошибки
			}
		}
		c.JSON(http.StatusOK, response)
		return
	case <-timeout:
		c.JSON(http.StatusRequestTimeout, gin.H{"Error": "Request timed out after 10 seconds"})
		return
	}
}

// GetUrlByShortName godoc
// @Summary Get original URL by short name
// @Description Retrieve the original URL by providing the short URL in the request
// @Tags URLs
// @Accept  json
// @Produce  json
// @Param link path string true "Short URL"
// @Success 200 {object} responses.GetUrlByShortNameResponse "Successfully retrieved the original URL"
// @Failure 400 {object} string "Invalid link"
// @Failure 404 {object} string "URL not found"
// @Failure 408 {object} string "Request timed out"
// @Failure 500 {object} string "Internal Server Error"
// @Router /url/{link} [get]
func (query *URLQueries) GetUrlByShortName(c *gin.Context) {
	link, exists := c.Get("link")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid link"})
		return
	}
	responseChan := make(chan responses.GetUrlByShortNameResponse)
	timeout := time.After(10 * time.Second)

	go func() {
		responseChan <- query.URLService.GetUrlByShortName(requests.GetUrlByShortNameRequest{
			ShortName: link.(string),
		})
	}()

	select {
	case response := <-responseChan:
		if response.Error != nil {
			if strings.Contains(response.Error.Error(), "url not found") {
				c.JSON(http.StatusNotFound, gin.H{"Error": response.Error.Error()})
				return // 404 если url не найден
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"Error": response.Error.Error()})
				return // Другие ошибки
			}
		}
		c.JSON(http.StatusOK, response)
		return
	case <-timeout:
		c.JSON(http.StatusRequestTimeout, gin.H{"Error": "Request timed out after 10 seconds"})
		return
	}
}

// DeleteByShortName godoc
// @Summary DeleteByShortName URL by short name
// @Description DeleteByShortName a URL by providing its short name in the request
// @Tags URLs
// @Accept  json
// @Produce  json
// @Param link path string true "Short URL to delete"
// @Success 200 {object} responses.DeleteUrlByShortNameResponse "Successfully deleted the URL"
// @Failure 400 {object} string "Invalid link"
// @Failure 404 {object} string "URL not found"
// @Failure 408 {object} string "Request timed out"
// @Failure 500 {object} string "Internal Server Error"
// @Router /url/{link} [delete]
func (query *URLQueries) DeleteByShortName(c *gin.Context) {
	link, exists := c.Get("link")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid link"})
		return
	}
	responseChan := make(chan responses.DeleteUrlByShortNameResponse)
	timeout := time.After(10 * time.Second)
	go func() {
		responseChan <- query.URLService.DeleteUrlByShortName(requests.DeleteByShortNameRequest{
			ShortName: link.(string),
		})
	}()

	select {
	case response := <-responseChan:
		if response.Error != nil {
			if strings.Contains(response.Error.Error(), "url not found") {
				c.JSON(http.StatusNotFound, gin.H{"Error": response.Error.Error()})
				return // 404 если url не найден
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"Error": response.Error.Error()})
				return // Другие ошибки
			}
		}
		c.JSON(http.StatusOK, response)
		return
	case <-timeout:
		c.JSON(http.StatusRequestTimeout, gin.H{"Error": "Request timed out after 10 seconds"})
		return
	}
}

// GetUrlStat godoc
// @Summary Get URL statistics by short name
// @Description Get statistics for a URL using its short name
// @Tags URLs
// @Accept  json
// @Produce  json
// @Param link path string true "Short URL to get statistics for"
// @Success 200 {object} responses.GetUrlStatByShortNameResponse "Successfully retrieved URL statistics"
// @Failure 400 {object} string "Invalid link"
// @Failure 404 {object} string "URL not found"
// @Failure 408 {object} string "Request timed out"
// @Failure 500 {object} string "Internal Server Error"
// @Router /url/stats/{link} [get]
func (query *URLQueries) GetUrlStat(c *gin.Context) {
	link, exists := c.Get("link")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid link"})
		return
	}
	responseChan := make(chan responses.GetUrlStatByShortNameResponse)
	timeout := time.After(10 * time.Second)
	go func() {
		responseChan <- query.URLService.GetUrlStatByShortName(requests.GetUrlStatByShortNameRequest{
			ShortName: link.(string),
		})
	}()

	select {
	case response := <-responseChan:
		if response.Error != nil {
			if strings.Contains(response.Error.Error(), "url not found") {
				c.JSON(http.StatusNotFound, gin.H{"Error": response.Error.Error()})
				return // 404 если url не найден
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"Error": response.Error.Error()})
				return // Другие ошибки
			}
		}
		c.JSON(http.StatusOK, response)
		return
	case <-timeout:
		c.JSON(http.StatusRequestTimeout, gin.H{"Error": "Request timed out after 10 seconds"})
		return
	}
}
