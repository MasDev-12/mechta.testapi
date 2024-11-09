package servers

import (
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

type MockRestServer struct {
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

func NewMockRestServer(ServerSettings *config.ServerSetting,
	DbSettings *config.DbSetting,
	Argon2Settings *config.Argon2Setting) *MockRestServer {
	server := &MockRestServer{
		router:         gin.Default(),
		ServerSettings: ServerSettings,
		DbSettings:     DbSettings,
		Argon2Settings: Argon2Settings,
	}
	server.Init() // Инициализация всех компонентов
	return server
}

func (s *MockRestServer) Init() {
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

func (s *MockRestServer) AddRoutes() {

	s.router.POST("/user/create", s.UserValidator.CreateUser(), s.UserCommands.CreateUserCommandExecute)
	s.router.GET("/user/:id", s.UserValidator.UserExists(), s.UserQueries.GetUserByIdQuery)
	s.router.POST("/url/shortener", s.UrlValidator.ValidateUrlForDuplicate(), s.URLCommands.CreateUrlCommandExecute)
	s.router.GET("/url/shortener/:userId", s.UrlValidator.ValidateUserExistsForTakeOwnUrls(), s.URLQueries.GetUserUrls)
	s.router.GET("/url/:link", s.UrlValidator.ShortUrlExists(), s.URLQueries.GetUrlByShortName)
	s.router.DELETE("/url/:link", s.UrlValidator.ShortUrlExists(), s.URLCommands.DeleteByShortName)
	s.router.GET("/url/stats/:link", s.UrlValidator.ShortUrlExists(), s.URLQueries.GetUrlStat)
}

func (s *MockRestServer) Router() *gin.Engine {
	return s.router
}

func (s *MockRestServer) StartMockServer() error {
	return nil
}
