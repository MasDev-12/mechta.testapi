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

func (query *URLQueries) Delete(c *gin.Context) {
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
