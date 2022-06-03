package modules

import (
	"gym-app/api"
	"gym-app/api/v1/auth"
	"gym-app/api/v1/user"
	authService "gym-app/business/user"
	jwtService "gym-app/business/user"
	userService "gym-app/business/user"
	"gym-app/config"
	userRepo "gym-app/repository/user"
	"gym-app/util"
)

func RegisterModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	userRepo := userRepo.RepositoryFactory(dbCon)

	userService := userService.NewUserService(userRepo)
	authService := authService.NewAuthService(userRepo)
	jwtService := jwtService.NewJWTService()

	controller := api.Controller{
		Auth: auth.NewAuthController(authService, jwtService, userService),
		User: user.NewUserController(userService, jwtService),
	}
	return controller
}
