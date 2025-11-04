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

	userGroup := r.Group("/user").Use(middleware.AuthMiddleware())
	{
		userGroup.GET("/:id", h.GetUserByID, middleware.ValidateUUID())
		userGroup.GET("/all", h.ListUser)
		userGroup.POST("", h.CreateUser, middleware.AdminMiddleware())                                  // добавлять может только админ
		userGroup.PUT("", h.UpdateUser, middleware.AdminMiddleware())                                   // изменять может только админ
		userGroup.DELETE("/:id", h.DeleteUser, middleware.ValidateUUID(), middleware.AdminMiddleware()) // удалять может только админ
	}
	r.GET("/user/email/:email", h.GetUserByEmail)
}
func initGroupModule(r *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewGroupRepository(db)
	sc := service.NewGroupService(repo)
	h := http.NewGroupHandler(sc)

	groupGroup := r.Group("/group").Use(middleware.AuthMiddleware()).Use(middleware.AdminMiddleware())
	{
		groupGroup.GET("/:id", h.GetByID, middleware.ValidateUUID())
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
		authGroup.GET("/logout", h.Logout)
		authGroup.GET("/status", h.AuthStatus)
	}
}
