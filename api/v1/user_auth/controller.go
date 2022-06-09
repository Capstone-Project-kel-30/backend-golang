package auth

import (
	"net/http"
	"strconv"

	"gym-app/api/common/response"
	service "gym-app/business/user"
	"gym-app/business/user/dto"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authService service.AuthService
	jwtService  service.JWTService
	userService service.UserService
}

func NewAuthController(
	authService service.AuthService,
	jwtService service.JWTService,
	userService service.UserService,
) *AuthController {
	return &AuthController{
		authService: authService,
		jwtService:  jwtService,
		userService: userService,
	}
}

func (controller *AuthController) Login(c echo.Context) error {
	var loginRequest dto.LoginRequest
	err := c.Bind(&loginRequest)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	err = controller.authService.VerifyCredential(loginRequest.Email, loginRequest.Password)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", "Invalid credentials", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	user, _ := controller.userService.FindUserByEmail(loginRequest.Email)

	token := controller.jwtService.GenerateToken((strconv.Itoa(user.ID)))
	user.Token = token
	response := response.BuildResponse(true, "ok", user)
	return c.JSON(http.StatusOK, response)
}

func (controller *AuthController) RegisterHandler(c echo.Context) error {
	var registerRequest dto.RegisterRequest
	err := c.Bind(&registerRequest)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", "Invalid request body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	user, err := controller.userService.CreateUser(registerRequest)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}
	// get otp
	totp := controller.authService.GenerateTOTP(user.Email)

	user.Totp = totp
	// send otp to email
	sendEmail := controller.authService.SendOTPtoEmail(totp, user.Name, user.Email)
	if sendEmail != nil {
		response := response.BuildErrorResponse("Failed to process request", sendEmail.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	token := controller.jwtService.GenerateToken((strconv.Itoa(user.ID)))
	user.Token = token
	response := response.BuildResponse(true, "OK", user)
	return c.JSON(http.StatusOK, response)
}
