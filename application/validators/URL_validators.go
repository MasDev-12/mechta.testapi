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

type URLValidator struct {
	URLService  *services.URLService
	UserService *services.UserService
}

func NewURLValidators(urlService *services.URLService,
	userService *services.UserService) *URLValidator {
	return &URLValidator{
		URLService:  urlService,
		UserService: userService}
}

func (urlValidator *URLValidator) UrlExists() gin.HandlerFunc {
	return func(c *gin.Context) {
		var url requests.CreateURLRequest
		if err := c.ShouldBindJSON(&url); err != nil {
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

		userExists := urlValidator.UserService.GetById(requests.GetUserRequest{
			Id: url.UserId,
		})

		if userExists.Id != nil && *userExists.Id != url.UserId {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user does not exist"})
			c.Abort()
			return
		}

		urlExists := urlValidator.URLService.GetUrlByOriginalName(requests.GetUrlByOriginalNameRequest{
			OriginalName: url.OriginalURL,
		})

		if urlExists.OriginalURL != nil && strings.ToLower(*urlExists.OriginalURL) == strings.ToLower(url.OriginalURL) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "url already exists"})
			c.Abort()
			return
		}

		c.Set("url", url)
		c.Next()
	}
}

func (urlValidator *URLValidator) UserExists() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIdParam := c.Param("userId")
		userId, err := uuid.Parse(userIdParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid user Id"})
			c.Abort()
			return
		}

		userExists := urlValidator.UserService.GetById(requests.GetUserRequest{
			Id: userId,
		})

		if userExists.Id != nil && *userExists.Id != userId {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user does not exist"})
			c.Abort()
			return
		}

		c.Set("userId", userId)
		c.Next()
	}
}

func (urlValidator *URLValidator) ShortUrlExists() gin.HandlerFunc {
	return func(c *gin.Context) {
		linkParam := c.Param("link")
		link := strings.TrimSpace(strings.ToLower(linkParam))
		if link == "" {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid shortener name"})
			c.Abort()
		}

		shortenerExists := urlValidator.URLService.GetUrlByShortName(requests.GetUrlByShortNameRequest{
			ShortName: link,
		})

		if shortenerExists.Url != nil && strings.ToLower(shortenerExists.Url.ShortURL) != strings.ToLower(link) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "shortUrl not exists"})
			c.Abort()
			return
		}

		c.Set("link", link)
		c.Next()
	}
}
