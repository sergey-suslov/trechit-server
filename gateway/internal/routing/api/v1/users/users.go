package users

import (
	"github.com/labstack/echo"
	models2 "github.com/sergey-suslov/trechit-server/gateway/internal/db/models"
	utils2 "github.com/sergey-suslov/trechit-server/gateway/utils"
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

	if err := utils2.Validate.Struct(form); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Wrong credentials")
	}

	existingUser, err := models2.GetUserByEmail(form.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error signing up")
	}
	if existingUser != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Email already exists")
	}

	if err := models2.CreateUser(form.Email, form.Name, form.Password); err != nil {
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

	if err := utils2.Validate.Struct(form); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Wrong body")
	}

	var user *models2.User
	user, err := models2.GetUserByEmail(form.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error getting user")
	}

	if user == nil || !models2.VerifyPassword(user, form.Password) {
		return echo.NewHTTPError(http.StatusBadRequest, "Wrong email or password")
	}

	expirationTime := time.Now().Add(5 * 24 * time.Hour)

	token, err := user.GetJwt(expirationTime)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error signing token")
	}

	e.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   *token,
		Expires: expirationTime,
	})

	return e.String(http.StatusOK, *token)
}

// GetProfile returns user model
func GetProfile(e echo.Context) error {
	user := e.Get("user").(*models2.UserClaims)
	return e.JSON(http.StatusOK, user)
}

// RefreshToken refresh user token
func RefreshToken(e echo.Context) error {
	userState := e.Get("user").(*models2.UserClaims)
	user, err := models2.GetUserByEmail(userState.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error refreshing token")
	}
	if user == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "No user for this token")
	}

	expirationTime := time.Now().Add(5 * 24 * time.Hour)
	token, err := user.GetJwt(expirationTime)

	if token == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error refreshing token")
	}

	e.SetCookie(&http.Cookie{
		Name:    "token",
		Value:   *token,
		Expires: expirationTime,
	})

	return e.String(http.StatusOK, *token)
}
