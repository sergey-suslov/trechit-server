package models

import (
	"crypto/sha256"
	"encoding/base64"
	"log"
	"math/rand"
	"time"
)

type User struct {
	id string
	email string
	hash string
	salt string
	name string
	created time.Time
}

// CreateUser create user model and write it to the DB
func CreateUser(email, name, password string) {
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