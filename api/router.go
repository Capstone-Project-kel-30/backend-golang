package api

import (
	"gym-app/api/middleware"
	"gym-app/api/v1/auth"
	"gym-app/api/v1/user"

	service "gym-app/business/user"

	"github.com/labstack/echo/v4"
)

var jwtService service.JWTService = service.NewJWTService()

type Controller struct {
	Auth *auth.AuthController
	User *user.UserController
}

func RegisterRoutes(e *echo.Echo, controller *Controller) {

	authRoutes := e.Group("/api/v1/auth")
	authRoutes.POST("/login", controller.Auth.Login)
	authRoutes.POST("/register", controller.Auth.RegisterHandler)

	userRoutes := e.Group("/api/v1/user", middleware.AuthorizeJWT(jwtService))
	userRoutes.GET("/profile", controller.User.Profile)
	userRoutes.PUT("/profile", controller.User.Update)

}
