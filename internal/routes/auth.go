package routes

import (
	"database/sql"
	"go-mitra-akosta/internal/delivery/http"
	"go-mitra-akosta/pkg/utils/config"

	"github.com/labstack/echo/v4"
)

func InitRouteAuth(Cfg *config.Config, db *sql.DB, route *echo.Echo) {
	// tokenAuthService := middleware.JWTAuthService(Cfg)
	controller := http.NewAuthController()

	apiv1 := route.Group("/api/v1")
	{
		auth := apiv1.Group("/auth")
		{
			auth.GET("/login", controller.SignInController)
			auth.GET("/register", controller.SignUpController)
		}

	}
}

// func Greetings(c echo.Context) error {
// 	return c.JSON(200, "Helloword")
// }
