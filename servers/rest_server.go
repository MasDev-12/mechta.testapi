package servers

import (
	"fmt"
	"github.com/MasDev-12/mechta.testapi/application/CQRS/commands"
	"github.com/MasDev-12/mechta.testapi/application/CQRS/queries"
	"github.com/MasDev-12/mechta.testapi/application/helpers"
	"github.com/MasDev-12/mechta.testapi/application/services"
	"github.com/MasDev-12/mechta.testapi/application/validators"
	"github.com/MasDev-12/mechta.testapi/config"
	"github.com/MasDev-12/mechta.testapi/docs"
	"github.com/MasDev-12/mechta.testapi/infrastructure/db_context"
	"github.com/MasDev-12/mechta.testapi/infrastructure/repositories"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type RestServer struct {
	router          *gin.Engine
	UserCommands    *commands.UserCommand
	UserQueries     *queries.UserQueries
	URLCommands     *commands.URLCommand
	URLQueries      *queries.URLQueries
	ServerSettings  *config.ServerSetting
	DbSettings      *config.DbSetting
	Argon2Settings  *config.Argon2Setting
	SwaggerSettings *config.SwaggerSetting
	UserValidator   *validators.UserValidator
	UrlValidator    *validators.URLValidator
}

func NewRestServer(ServerSettings *config.ServerSetting,
	DbSettings *config.DbSetting,
	Argon2Settings *config.Argon2Setting,
	swaggerSettings *config.SwaggerSetting) *RestServer {
	server := &RestServer{
		router:          gin.Default(),
		ServerSettings:  ServerSettings,
		DbSettings:      DbSettings,
		Argon2Settings:  Argon2Settings,
		SwaggerSettings: swaggerSettings,
	}
	server.Init() // Инициализация всех компонентов
	return server
}

func (s *RestServer) Init() {
	// Инициализация DbContext
	dbContext := db_context.NewDbContext(s.DbSettings)

	// Инициализация репозиториев
	userRepository := repositories.NewUserRepository(dbContext)
	urlRepository := repositories.NewURLRepository(dbContext)

	//Инициализация хэлперов
	argon2Helper := helpers.NewArgon2Helper(s.Argon2Settings)
	// Инициализация сервисов

	userService := services.NewUserService(userRepository, argon2Helper)
	urlService := services.NewURLService(urlRepository)

	// Инициализация команд
	s.UserCommands = commands.NewUserCommand(userService)
	s.UserQueries = queries.NewUserQueries(userService)
	s.URLCommands = commands.NewURLCommand(urlService)
	s.URLQueries = queries.NewURLQueries(urlService)

	s.UserValidator = validators.NewUserValidator(userRepository)
	s.UrlValidator = validators.NewURLValidators(urlRepository, userRepository)
	s.AddRoutes()
}

func (s *RestServer) AddRoutes() {
	docs.SwaggerInfo.Host = s.SwaggerSettings.Host
	docs.SwaggerInfo.Description = s.SwaggerSettings.Description
	docs.SwaggerInfo.Title = s.SwaggerSettings.PageTitle
	docs.SwaggerInfo.Version = s.SwaggerSettings.Version
	docs.SwaggerInfo.BasePath = s.SwaggerSettings.BasePath

	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.router.POST("/api/v1/user/create", s.UserValidator.CreateUser(), s.UserCommands.CreateUserCommandExecute)
	s.router.GET("/api/v1/user/:id", s.UserValidator.UserExists(), s.UserQueries.GetUserByIdQuery)
	s.router.POST("/api/v1/url/shortener", s.UrlValidator.ValidateUrlForDuplicate(), s.URLCommands.CreateUrlCommandExecute)
	s.router.GET("/api/v1/url/shortener/:userId", s.UrlValidator.ValidateUserExistsForTakeOwnUrls(), s.URLQueries.GetUserUrls)
	s.router.GET("/api/v1/url/:link", s.UrlValidator.ShortUrlExists(), s.URLQueries.GetUrlByShortName)
	s.router.DELETE("/api/v1/url/:link", s.UrlValidator.ShortUrlExists(), s.URLCommands.DeleteByShortName)
	s.router.GET("/api/v1/url/stats/:link", s.UrlValidator.ShortUrlExists(), s.URLQueries.GetUrlStat)
}

func (s *RestServer) Start() error {
	port := fmt.Sprintf("%s:%d", s.ServerSettings.Host, s.ServerSettings.Port)
	return s.router.Run(port)
}

func (s *RestServer) Router() *gin.Engine {
	return s.router
}
