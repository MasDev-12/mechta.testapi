package commands

import (
	"github.com/MasDev-12/mechta.testapi/application/CQRS/requests"
	"github.com/MasDev-12/mechta.testapi/application/CQRS/responses"
	"github.com/MasDev-12/mechta.testapi/application/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserCommand struct {
	UserService *services.UserService
}

func NewUserCommand(userService *services.UserService) *UserCommand {
	return &UserCommand{
		UserService: userService,
	}
}

func (command *UserCommand) CreateCommandExecute(c *gin.Context) {
	var request requests.CreateUserRequest
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusBadRequest, "Ошибка: неверные данные пользователя")
		return
	}
	request = user.(requests.CreateUserRequest)

	responseChan := make(chan responses.CreateUserResponse)
	timeout := time.After(10 * time.Second)
	go func() {
		responseChan <- command.UserService.Create(request)
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
