package models

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx/v4"
	"github.com/sergey-suslov/trechit-server/gateway/internal/db"
	"math/rand"
	"os"
	"time"
)

type User struct {
	Id      int32
	Email   string
	Hash    string
	Salt    string
	Name    string
	Created time.Time
}

type UserClaims struct {
	Id    int32  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.StandardClaims
}

// GetJwt returns jwt encoded string for current user
func (user *User) GetJwt(experationTime time.Time) (*string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	claims := &UserClaims{
		Id:    user.Id,
		Email: user.Email,
		Name:  user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: experationTime.Unix(),
		},
	}
	tokenBase := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenBase.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, err
	}

	return &token, err
}

// GenerateHash returns user password hash
func GenerateHash(password, salt string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password + salt))
	hash := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return hash
}

// GenerateHashWithRandomSalt returns user password hash with random salt
func GenerateHashWithRandomSalt(password string) (string, string) {
	salt := getSalt(80)
	hash := GenerateHash(password, salt)
	return hash, salt
}

// CreateUser create user model and write it to the DB
func CreateUser(email, name, password string) error {
	hash, salt := GenerateHashWithRandomSalt(password)

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

// GetUserByEmail returns user or error
func GetUserByEmail(email string) (*User, error) {
	var user User
	err := db.Pool.QueryRow(context.Background(), "select id, email, name, created, hash, salt from users where Email=$1", email).Scan(&user.Id, &user.Email, &user.Name, &user.Created, &user.Hash, &user.Salt)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// VerifyPassword compares user password and given password
func VerifyPassword(user *User, password string) bool {
	return user.Hash == GenerateHash(password, user.Salt)
}
