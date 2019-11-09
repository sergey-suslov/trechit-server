package users

import (
	"github.com/sergey-suslov/trechit-server/internal/db/models"
	"testing"
)

func generateUserModel(password, salt string) *models.User {
	hash, salt := models.GenerateHashWithRandomSalt(password)
	user := &models.User{
		Id:    1,
		Name:  "Test",
		Email: "Email",
		Hash:  hash,
		Salt:  salt,
	}
	return user
}

func TestVerifyPassword(t *testing.T) {
	pass := "testpass"
	salt := "testsalt"
	user := generateUserModel(pass, salt)

	t.Run("Passwords match", func(t *testing.T) {
		if !models.VerifyPassword(user, pass) {
			t.Error("Passwords match, should return true")
		}
	})
	t.Run("Passwords do not match", func(t *testing.T) {
		if models.VerifyPassword(user, pass + "1") {
			t.Error("Passwords do not match, should return false")
		}
	})
}
