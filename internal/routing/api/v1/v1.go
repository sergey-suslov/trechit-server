package v1

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sergey-suslov/trechit-server/internal/db/models"
	"github.com/sergey-suslov/trechit-server/internal/routing/api/v1/users"
	"os"
)

// InitAPI initialize api v1
func InitAPI(e *echo.Group) {
	e.POST("/sign-up", users.SignUp)
	e.POST("/sign-in", users.Auth)
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(os.Getenv("JWT_SECRET")),
		TokenLookup: "cookie:token",
		Claims:      &models.UserClaims{},
	}))
	e.GET("/profile", users.GetProfile)
}
