package api

import (
	"gym-app/api/middleware"
	"gym-app/api/v1/user"
	user_auth "gym-app/api/v1/user_auth"

	service "gym-app/business/user"

	"github.com/labstack/echo/v4"
)

var jwtService service.JWTService = service.NewJWTService()

type Controller struct {
	UserAuth *user_auth.AuthController
	User     *user.UserController
}

func RegisterRoutes(e *echo.Echo, controller *Controller) {

	userAuthRoutes := e.Group("/api/v1/auth")
	userAuthRoutes.POST("/login", controller.UserAuth.Login)
	userAuthRoutes.POST("/register", controller.UserAuth.RegisterHandler)

	userRoutes := e.Group("/api/v1/user", middleware.AuthorizeJWT(jwtService))
	userRoutes.GET("/profile", controller.User.Profile)
	userRoutes.PUT("/profile", controller.User.Update)

}
