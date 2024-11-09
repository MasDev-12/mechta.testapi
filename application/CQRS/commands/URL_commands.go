package commands

import (
	"github.com/MasDev-12/mechta.testapi/application/CQRS/requests"
	"github.com/MasDev-12/mechta.testapi/application/CQRS/responses"
	"github.com/MasDev-12/mechta.testapi/application/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type URLCommand struct {
	URLService *services.URLService
}

func NewURLCommand(urlService *services.URLService) *URLCommand {
	return &URLCommand{
		URLService: urlService,
	}
}

// CreateUrl godoc
// @Summary Create a new shortened URL
// @Description Create a shortened URL by passing the URL data in the request body
// @Tags URLs
// @Accept  json
// @Produce  json
// @Param url body requests.CreateURLRequest true "URL data"
// @Success 200 {object} responses.CreateUrlResponse "URL successfully shortened"
// @Failure 400 {object} string "Invalid URL data"
// @Failure 408 {object} string "Request timed out"
// @Failure 500 {object} string "Internal Server Error"
// @Router /url/shortener [post]
func (command *URLCommand) CreateUrlCommandExecute(c *gin.Context) {
	var request requests.CreateURLRequest
	url, exists := c.Get("url")
	if !exists {
		c.JSON(http.StatusBadRequest, "Error: error data of url")
		return
	}
	request = url.(requests.CreateURLRequest)

	responseChan := make(chan responses.CreateUrlResponse)
	timeout := time.After(10 * time.Second)
	go func() {
		responseChan <- command.URLService.Create(request)
	}()
	select {
	case response := <-responseChan:
		if response.Error != nil {
			c.JSON(http.StatusBadRequest, response.Error)
			return
		} else {
			c.JSON(http.StatusOK, response)
			return
		}
	case <-timeout:
		c.JSON(http.StatusRequestTimeout, gin.H{"Error": "Request timed out after 10 seconds"})
		return
	}
}
