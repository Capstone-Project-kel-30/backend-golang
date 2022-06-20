package api

import (
	"gym-app/api/middleware"
	"gym-app/api/v1/auth"
	"gym-app/api/v1/member"
	"gym-app/api/v1/user"

	service "gym-app/business/user"

	"github.com/labstack/echo/v4"
)

var jwtService service.JWTService = service.NewJWTService()

type Controller struct {
	UserAuth *auth.AuthController
	User     *user.UserController
	Member   *member.MemberController
}

func RegisterRoutes(e *echo.Echo, controller *Controller) {

	userAuthRoutes := e.Group("/api/v1/auth")
	userAuthRoutes.POST("/login", controller.UserAuth.Login)
	userAuthRoutes.POST("/register", controller.UserAuth.RegisterHandler)
	userAuthRoutes.POST("/verify-email", controller.UserAuth.SendEmailVerification)
	userAuthRoutes.POST("/reset-password", controller.UserAuth.ResetPassword)

	userRoutes := e.Group("/api/v1/user", middleware.AuthorizeJWT(jwtService))
	userRoutes.GET("/profile", controller.User.Profile)
	userRoutes.PUT("/profile", controller.User.Update)

	memberRoutes := e.Group("/api/v1/member", middleware.AuthorizeJWT(jwtService))
	memberRoutes.POST("/register", controller.Member.RegisterMember)
}
