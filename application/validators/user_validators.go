package validators

import (
	"errors"
	"github.com/MasDev-12/mechta.testapi/application/CQRS/requests"
	"github.com/MasDev-12/mechta.testapi/application/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

type UserValidator struct {
	UserService *services.UserService
}

func NewUserValidator(userService *services.UserService) *UserValidator {
	return &UserValidator{UserService: userService}
}

func (userValidator *UserValidator) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user requests.CreateUserRequest
		if err := c.ShouldBindJSON(&user); err != nil {
			var validationErrors validator.ValidationErrors
			exists := errors.As(err, &validationErrors)
			if exists == true {
				errorsMap := make(map[string]string)
				for _, fieldError := range validationErrors {
					errorsMap[fieldError.Field()] = fieldError.Error()
				}
				c.JSON(http.StatusBadRequest, gin.H{"error": errorsMap})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			c.Abort()
			return
		}

		emailExists := userValidator.UserService.GetByEmail(requests.GetUserByEmailRequest{
			Email: user.Email,
		})

		if emailExists.Email != nil && strings.ToLower(*emailExists.Email) == strings.ToLower(user.Email) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exist"})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func (userValidator *UserValidator) UserExists() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid user Id"})
			c.Abort()
			return
		}

		userExists := userValidator.UserService.GetById(requests.GetUserRequest{
			Id: id,
		})

		if userExists.Id != nil && userExists.Id == nil || *userExists.Id != id {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not exists"})
			c.Abort()
			return
		}

		c.Set("userId", id)
		c.Next()
	}
}
