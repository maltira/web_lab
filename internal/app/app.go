package app

import (
	"fmt"
	"web-lab/configs"
	"web-lab/internal/delivery/http"
	"web-lab/internal/middleware"
	"web-lab/internal/repository"
	"web-lab/internal/service"
	"web-lab/pkg/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Run() {
	// ? Инициализация конфига
	config.Init()

	// ? База данных
	db := database.InitDatabase()
	database.Migrate(db)

	// ? Инициализация модулей и путей
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	apiGroup := router.Group("/api/v1")

	// ? Инициализаця модуля
	initUserModule(apiGroup, db)
	initGroupModule(apiGroup, db)
	initAuthModule(apiGroup, db)
	initPublicationModule(apiGroup, db)

	// ? Запуск сервера
	fmt.Printf("Запуск сервера по адресу: http://%s:%s/api/v1\n", config.Cfg.AppHost, config.Cfg.AppPort)
	err := router.Run(":" + config.Cfg.AppPort)
	if err != nil {
		panic(fmt.Sprintf("Не удалось запустить сервер: %s", err))
	}
}

func initUserModule(r *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewUserRepository(db)
	sc := service.NewUserService(repo)
	h := http.NewUserHandler(sc)

	userGroup := r.Group("/user")
	{
		userGroup.GET("/:id", middleware.ValidateUUID(), h.GetUserByID)
		userGroup.GET("/all", h.ListUser)
		userGroup.GET("", middleware.AuthMiddleware(), h.GetCurrentUser)
		userGroup.POST("", h.CreateUser)
		userGroup.PUT("", h.UpdateUser)
		userGroup.PUT("/greeting", h.UpdateUserGreeting)
		userGroup.DELETE("/:id", middleware.ValidateUUID(), h.DeleteUser)
	}
	r.GET("/user/email/:email", h.GetUserByEmail)
}
func initGroupModule(r *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewGroupRepository(db)
	sc := service.NewGroupService(repo)
	h := http.NewGroupHandler(sc)

	groupGroup := r.Group("/group").Use(middleware.AuthMiddleware()).Use(middleware.AdminMiddleware())
	{
		groupGroup.GET("/:id", middleware.ValidateUUID(), h.GetByID)
		groupGroup.GET("/all", h.GetAll)
	}
}
func initAuthModule(r *gin.RouterGroup, db *gorm.DB) {
	repoGroup := repository.NewGroupRepository(db)
	repoUser := repository.NewUserRepository(db)
	scGroup := service.NewGroupService(repoGroup)
	scUser := service.NewUserService(repoUser)
	h := http.NewAuthHandler(scUser, scGroup)

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", h.Login)
		authGroup.POST("/registration", h.Registration)
		authGroup.POST("/logout", h.Logout)
		authGroup.GET("/status", h.AuthStatus)
	}
}
func initPublicationModule(r *gin.RouterGroup, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userSc := service.NewUserService(userRepo)
	repo := repository.NewPublicationRepository(db)
	sc := service.NewPublicationService(repo)
	h := http.NewPublicationHandler(sc, userSc)

	publicationGroup := r.Group("/publication").Use(middleware.AuthMiddleware())
	{
		publicationGroup.POST("/create", h.CreatePublication)
		publicationGroup.PUT("/update", h.UpdatePublication)
		publicationGroup.DELETE("/:id", middleware.ValidateUUID(), h.DeletePublication)
		publicationGroup.GET("/:id", middleware.ValidateUUID(), h.FindByID)
	}
	r.GET("/publication/all", h.FindAllPublications)
}
