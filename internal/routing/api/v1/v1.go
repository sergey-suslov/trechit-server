package v1

import (
	"github.com/labstack/echo"
	"github.com/sergey-suslov/trechit-server/internal/routing/api/v1/users"
)

// InitAPI initialize api v1
func InitAPI(e *echo.Group) {
	e.POST("/sign-up", users.SignUp)
}
