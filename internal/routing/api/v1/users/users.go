package users

import (
	"github.com/labstack/echo"
	"github.com/sergey-suslov/trechit-server/internal/db/models"
	"github.com/sergey-suslov/trechit-server/utils"
	"net/http"
)

// CreateUserForm request body for user signup
type CreateUserForm struct {
	Email    string `form:"email" validate:"required,email"`
	Name     string `form:"name" validate:"required"`
	Password string `form:"password" validate:"required"`
}

// SignUp sign up controller
func SignUp(c echo.Context) error {
	var form CreateUserForm
	if err := c.Bind(&form); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Wrong credentials")
	}

	if err := utils.Validate.Struct(form); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Wrong credentials")
	}

	models.CreateUser(form.Email, form.Name, form.Password)

	return c.String(200, "")
}
