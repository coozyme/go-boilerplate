package http

import (
	"go-mitra-akosta/pkg/utils/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (a *AuthController) SignInController(c echo.Context) error {

	return response.ResponseSuccess(http.StatusFound, 200, "success", map[string]string{
		"name":    "Ary",
		"address": "Tangerang",
	}).Send(c)
}

func (a *AuthController) SignUpController(c echo.Context) error {

	return response.ResponsepSuccessPagination(http.StatusFound, 200, "mantap", "okedehß", &response.Pagination{
		Count:          500,
		Limit:          10,
		PaginationInfo: response.PaginationInfo{Page: 1, TotalPage: 100, PerPage: 2, CurrentPage: 3},
	}).Send(c)
}
