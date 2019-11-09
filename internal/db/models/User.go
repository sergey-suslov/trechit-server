package models

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"github.com/sergey-suslov/trechit-server/internal/db"
	"log"
	"math/rand"
	"time"
)

type User struct {
	id      string
	email   string
	hash    string
	salt    string
	name    string
	created time.Time
}

// CreateUser create user model and write it to the DB
func CreateUser(email, name, password string) error {
	salt := getSalt(80)
	hasher := sha256.New()
	hasher.Write([]byte(password + salt))
	hash := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	log.Printf(`
		New User---
		Email: %s
		Name: %s
		Salt %s
		Hash %s
	`, email, name, salt, hash)
	// TODO create user

	_, err := db.Pool.Exec(context.Background(), `
		insert into users values(default, $1, $2, $3, $4, default);
	`, name, email, hash, salt)
	if err != nil {
		return err
	}
	return nil
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func getSalt(length int) string {
	return stringWithCharset(length, charset)
}
