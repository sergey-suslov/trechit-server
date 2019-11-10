package users

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/sergey-suslov/trechit-server/internal/db/models"
	"github.com/sergey-suslov/trechit-server/utils"
	"net/http"
	"time"
)

// CreateUserForm request body for user signup
type CreateUserForm struct {
	Email    string `form:"email" validate:"required,email"`
	Name     string `form:"name" validate:"required"`
	Password string `form:"password" validate:"required"`
}

type AuthUserForm struct {
	Email    string `form:"email" validate:"required,email"`
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

	existingUser, err := models.GetUserByEmail(form.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error signing up")
	}
	if existingUser != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Email already exists")
	}

	if err := models.CreateUser(form.Email, form.Name, form.Password); err != nil {
	    return echo.NewHTTPError(http.StatusInternalServerError, "Error creating user")
	}

	return c.String(200, "")
}

// Auth Authorize user
func Auth(e echo.Context) error {
	var form AuthUserForm
	if err := e.Bind(&form); err != nil {
	    return echo.NewHTTPError(http.StatusBadRequest, "Wrong body")
	}

	if err := utils.Validate.Struct(form); err != nil {
	    return echo.NewHTTPError(http.StatusBadRequest, "Wrong body")
	}

	var user *models.User
	user, err := models.GetUserByEmail(form.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error getting user")
	}

	if user == nil || !models.VerifyPassword(user, form.Password) {
		return echo.NewHTTPError(http.StatusBadRequest, "Wrong email or password")
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	token, err := user.GetJwt(expirationTime)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error signing token")
	}

	e.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   *token,
		Expires: expirationTime,
	})

	return e.String(http.StatusOK, "")
}

// GetProfile returns user model
func GetProfile(e echo.Context) error {
	token := e.Get("user").(*jwt.Token)
	user := token.Claims.(*models.UserClaims)
	return e.JSON(http.StatusOK, user)
}