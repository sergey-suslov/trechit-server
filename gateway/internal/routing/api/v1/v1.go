package v1

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	models2 "github.com/sergey-suslov/trechit-server/gateway/internal/db/models"
	socket2 "github.com/sergey-suslov/trechit-server/gateway/internal/routing/api/v1/socket"
	users2 "github.com/sergey-suslov/trechit-server/gateway/internal/routing/api/v1/users"
	"os"
)

// InitAPI initialize api v1
func InitAPI(e *echo.Group) {
	socketGroup := e.Group("/socket")
	socketGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(os.Getenv("JWT_SECRET")),
		TokenLookup: "query:token",
		Claims:      &models2.UserClaims{},
	}))
	socketGroup.GET("", socket2.InitSocketConn)

	e.POST("/sign-up", users2.SignUp)
	e.POST("/sign-in", users2.Auth)
	e.GET("/socket", socket2.InitSocketConn)
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(os.Getenv("JWT_SECRET")),
		TokenLookup: "cookie:token",
		Claims:      &models2.UserClaims{},
	}))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			user := token.Claims.(*models2.UserClaims)
			c.Set("user", user)
			return next(c)
		}
	})
	e.GET("/profile", users2.GetProfile)
	e.POST("/refresh", users2.RefreshToken)
}
