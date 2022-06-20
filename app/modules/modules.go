package modules

import (
	"gym-app/api"
	"gym-app/api/v1/auth"
	"gym-app/api/v1/member"
	"gym-app/api/v1/user"
	memberService "gym-app/business/member"
	authService "gym-app/business/user"
	jwtService "gym-app/business/user"
	userService "gym-app/business/user"
	"gym-app/config"
	memberRepo "gym-app/repository/member"
	userRepo "gym-app/repository/user"
	"gym-app/util"
)

func RegisterModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	userRepo := userRepo.RepositoryFactory(dbCon)

	userService := userService.NewUserService(userRepo)
	authService := authService.NewAuthService(userRepo)
	jwtService := jwtService.NewJWTService()

	memberRepo := memberRepo.RepositoryFactory(dbCon)
	memberService := memberService.NewMemberService(memberRepo)

	controller := api.Controller{
		UserAuth: auth.NewAuthController(authService, jwtService, userService),
		User:     user.NewUserController(userService, jwtService),
		Member:   member.NewMemberController(memberService, jwtService),
	}
	return controller
}
