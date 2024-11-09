package validators

import (
	"errors"
	"github.com/MasDev-12/mechta.testapi/application/CQRS/requests"
	"github.com/MasDev-12/mechta.testapi/infrastructure/repositories"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

type URLValidator struct {
	UrlRepository  *repositories.URLRepository
	UserRepository *repositories.UserRepository
}

func NewURLValidators(urlRepository *repositories.URLRepository,
	userRepository *repositories.UserRepository) *URLValidator {
	return &URLValidator{
		UrlRepository:  urlRepository,
		UserRepository: userRepository}
}

func (urlValidator *URLValidator) ValidateUrlForDuplicate() gin.HandlerFunc {
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

		_, userError := urlValidator.UserRepository.GetById(url.UserId)

		if userError != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": userError.Error()})
			c.Abort()
			return
		}

		urlExists, urlError := urlValidator.UrlRepository.GetUrlByOriginalName(url.OriginalURL)

		if urlError != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": urlError.Error()})
			c.Abort()
			return
		}

		if urlExists != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "url already exists"})
			c.Abort()
			return
		}

		c.Set("url", url)
		c.Next()
	}
}

func (urlValidator *URLValidator) ValidateUserExistsForTakeOwnUrls() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIdParam := c.Param("userId")
		userId, err := uuid.Parse(userIdParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid user Id"})
			c.Abort()
			return
		}

		_, userError := urlValidator.UserRepository.GetById(userId)

		if userError != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": userError.Error()})
			c.Abort()
			return
		}

		c.Set("userId", userId)
		c.Next()
	}
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

		_, userError := urlValidator.UserRepository.GetById(url.UserId)

		if userError != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": userError.Error()})
			c.Abort()
			return
		}

		urlExists, urlError := urlValidator.UrlRepository.GetUrlByOriginalName(url.OriginalURL)

		if urlError != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": urlError.Error()})
			c.Abort()
			return
		}

		if urlExists == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "url doesn't exists"})
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

		_, userError := urlValidator.UserRepository.GetById(userId)

		if userError != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": userError.Error()})
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
			return
		}

		shortenerExists, shortenerError := urlValidator.UrlRepository.GetUrlByShortName(link)

		if shortenerError != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "shortUrl not exists"})
			c.Abort()
			return
		}
		if shortenerExists == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "shortUrl not exists"})
			c.Abort()
			return
		}

		c.Set("link", link)
		c.Next()
	}
}
