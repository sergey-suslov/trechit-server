package users

import (
	"github.com/labstack/echo"
)

// SignUpForm request body for user signup
type SignUpForm struct {
	email    string `form:"email"`
	name     string `form:"name"`
	password string `form:"password"`
}

// SignUp sign up controller
func SignUp(c echo.Context) error {
	email := c.FormValue("email")
	return c.String(200, "Email: "+email)
}
