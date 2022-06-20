package member

import (
	"fmt"
	"gym-app/api/common/response"
	service "gym-app/business/member"
	"gym-app/business/member/dto"
	jwtService "gym-app/business/user"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type MemberController struct {
	memberService service.MemberService
	jwtService    jwtService.JWTService
}

func NewMemberController(
	memberService service.MemberService,
	jwtService jwtService.JWTService,
) *MemberController {
	return &MemberController{
		memberService: memberService,
		jwtService:    jwtService,
	}
}

func (controller *MemberController) RegisterMember(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	token := controller.jwtService.ValidateToken(authHeader, c)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	_ = userID

	var RegisterMemberRequest dto.RegisterMemberRequest
	if err := c.Bind(RegisterMemberRequest); err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	member, err := controller.memberService.CreateMember(RegisterMemberRequest)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := member
	return c.JSON(http.StatusOK, response)
}
