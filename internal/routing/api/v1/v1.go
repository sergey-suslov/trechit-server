package v1

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sergey-suslov/trechit-server/internal/db/models"
	"github.com/sergey-suslov/trechit-server/internal/routing/api/v1/socket"
	"github.com/sergey-suslov/trechit-server/internal/routing/api/v1/users"
	"os"
)

// InitAPI initialize api v1
func InitAPI(e *echo.Group) {
	socketGroup := e.Group("/socket")
	socketGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(os.Getenv("JWT_SECRET")),
		TokenLookup: "query:token",
		Claims:      &models.UserClaims{},
	}))
	socketGroup.GET("", socket.InitSocketConn)

	e.POST("/sign-up", users.SignUp)
	e.POST("/sign-in", users.Auth)
	e.GET("/socket", socket.InitSocketConn)
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(os.Getenv("JWT_SECRET")),
		TokenLookup: "cookie:token",
		Claims:      &models.UserClaims{},
	}))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			user := token.Claims.(*models.UserClaims)
			c.Set("user", user)
			return next(c)
		}
	})
	e.GET("/profile", users.GetProfile)
	e.POST("/refresh", users.RefreshToken)
}
