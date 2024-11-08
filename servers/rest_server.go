package servers

import (
	"fmt"
	"github.com/MasDev-12/mechta.testapi/application/CQRS/commands"
	"github.com/MasDev-12/mechta.testapi/application/CQRS/queries"
	"github.com/MasDev-12/mechta.testapi/application/helpers"
	"github.com/MasDev-12/mechta.testapi/application/services"
	"github.com/MasDev-12/mechta.testapi/application/validators"
	"github.com/MasDev-12/mechta.testapi/config"
	"github.com/MasDev-12/mechta.testapi/infrastructure/db_context"
	"github.com/MasDev-12/mechta.testapi/infrastructure/repositories"
	"github.com/gin-gonic/gin"
)

type RestServer struct {
	router         *gin.Engine
	UserCommands   *commands.UserCommand
	UserQueries    *queries.UserQueries
	URLCommands    *commands.URLCommand
	URLQueries     *queries.URLQueries
	ServerSettings *config.ServerSetting
	DbSettings     *config.DbSetting
	Argon2Settings *config.Argon2Setting
	UserValidator  *validators.UserValidator
	UrlValidator   *validators.URLValidator
}

func NewRestServer(ServerSettings *config.ServerSetting,
	DbSettings *config.DbSetting,
	Argon2Settings *config.Argon2Setting) *RestServer {
	server := &RestServer{
		router:         gin.Default(),
		ServerSettings: ServerSettings,
		DbSettings:     DbSettings,
		Argon2Settings: Argon2Settings,
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

	s.UserValidator = validators.NewUserValidator(userService)
	s.UrlValidator = validators.NewURLValidators(urlService, userService)
	s.AddRoutes()
}

func (s *RestServer) AddRoutes() {
	s.router.POST("/user/create", s.UserValidator.CreateUser(), s.UserCommands.CreateCommandExecute)
	s.router.GET("/user/:id", s.UserValidator.UserExists(), s.UserQueries.GetUserByIdQuery)
	s.router.POST("/url/shortener", s.UrlValidator.UrlExists(), s.URLCommands.CreateUrlCommandExecute)
	s.router.GET("/url/shortener/:userId", s.UrlValidator.UrlExists(), s.URLQueries.GetUserUrls)
	s.router.GET("/url/:link", s.UrlValidator.UrlExists(), s.URLQueries.GetUrlByShortName)
	s.router.DELETE("/url/:link", s.UrlValidator.UrlExists(), s.URLQueries.Delete)
	s.router.GET("/url/stats/:link", s.UrlValidator.UrlExists(), s.URLQueries.GetUrlStat)
}

func (s *RestServer) Start() error {
	port := fmt.Sprintf("%s:%d", s.ServerSettings.Host, s.ServerSettings.Port)
	return s.router.Run(port)
}

func (s *RestServer) Router() *gin.Engine {
	return s.router
}
