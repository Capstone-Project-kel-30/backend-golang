package user

import (
	"fmt"
	"net/http"
	"strconv"

	"gym-app/api/common/response"
	service "gym-app/business/user"
	"gym-app/business/user/dto"

	"github.com/golang-jwt/jwt"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(
	userService service.UserService,
	jwtService service.JWTService,
) *UserController {
	return &UserController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (controller *UserController) Profile(c echo.Context) error {
	header := c.Request().Header.Get("Authorization")
	token := controller.jwtService.ValidateToken(header, c)
	if header == "" {
		response := response.BuildErrorResponse("Failed to process request", "Failed to validate token", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user, err := controller.userService.FindUserByID(id)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := response.BuildResponse(true, "User found", user)
	return c.JSON(http.StatusOK, response)
}

func (controller *UserController) Update(c echo.Context) error {
	var updateUserRequest dto.UpdateUserRequest
	header := c.Request().Header.Get("Authorization")
	token := controller.jwtService.ValidateToken(header, c)
	if header == "" {
		response := response.BuildErrorResponse("Failed to process request", "Failed to validate token", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	if err := c.Bind(&updateUserRequest); err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	intID, err := strconv.Atoi(id)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	updateUserRequest.ID = intID
	user, err := controller.userService.UpdateUser(updateUserRequest)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	response := response.BuildResponse(true, "User updated", user)
	return c.JSON(http.StatusOK, response)

}

func (controller *UserController) getUserIDByHeader(c echo.Context) string {
	header := c.Request().Header.Get("Authorization")

	if header == "" {
		return fmt.Sprint("Error", "Failed to validate token")
	}
	token := controller.jwtService.ValidateToken(header, c)
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
