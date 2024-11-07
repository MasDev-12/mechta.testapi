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

type UserQueries struct {
	UserService *services.UserService
}

func NewUserQueries(userService *services.UserService) *UserQueries {
	return &UserQueries{
		UserService: userService,
	}
}

func (query *UserQueries) GetUserByIdQuery(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid user Id"})
		return
	}
	responseChan := make(chan responses.GetUserResponse)
	timeout := time.After(10 * time.Second)

	go func() {
		responseChan <- query.UserService.GetById(requests.GetUserRequest{Id: id.(uuid.UUID)})
	}()

	select {
	case response := <-responseChan:
		if response.Error != nil {
			if strings.Contains(response.Error.Error(), "user not found") {
				c.JSON(http.StatusNotFound, gin.H{"Error": response.Error.Error()})
				return // 404 если user не найден
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