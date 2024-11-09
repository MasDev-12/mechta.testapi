package validators

import (
	"errors"
	"github.com/MasDev-12/mechta.testapi/application/CQRS/requests"
	"github.com/MasDev-12/mechta.testapi/infrastructure/repositories"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
)

type UserValidator struct {
	UserRepository *repositories.UserRepository
}

func NewUserValidator(userRepository *repositories.UserRepository) *UserValidator {
	return &UserValidator{UserRepository: userRepository}
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

		_, userError := userValidator.UserRepository.GetUserByEmail(user.Email)
		if userError != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user already exists"})
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

		userExists, userError := userValidator.UserRepository.GetById(id)

		if userExists == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not exists"})
			c.Abort()
			return
		}

		if userError != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": userError.Error()})
			c.Abort()
			return
		}

		c.Set("userId", id)
		c.Next()
	}
}
