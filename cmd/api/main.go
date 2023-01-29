package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"utfpr_db/config"
	"utfpr_db/internal/controller"
	"utfpr_db/internal/controller/v1"
	"utfpr_db/internal/facade"
	"utfpr_db/internal/middleware"
	"utfpr_db/internal/repository"
	"utfpr_db/internal/service"
	"utfpr_db/rest"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @termsOfService http://swagger.io/terms/
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()

	db, err := gorm.Open(postgres.Open(cfg.DbConnString), &gorm.Config{FullSaveAssociations: true})
	if err != nil {
		fmt.Println("error opening connection ", err)
	}

	healthCheckController := controller.NewHealthCheckController()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userFacade := facade.NewUserFacade(userService)
	userController := v1.NewUserController(userFacade)

	securityFacade := facade.NewSecurityFacade()
	authMiddleware := middleware.NewAuthMiddleware(securityFacade)

	serverRest := rest.NewRestServer(
		cfg,
		&rest.Controllers{
			UserController:        userController,
			HealthCheckController: healthCheckController,
		},
		&rest.Middlewares{
			AuthMiddleware: authMiddleware,
		},
	)
	serverRest.StartListening()
}
