package types

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the system.
type User struct {
	ID           int    `json:"id,omitempty"`
	Username     string `json:"username,omitempty"`
	PasswordHash string `json:"password_hash,omitempty"`
}

type RequestUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUser(ru RequestUser) (User, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(ru.Password), bcrypt.DefaultCost)

	if err != nil {
		return User{}, err
	}

	return User{
		Username:     ru.Username,
		PasswordHash: string(hashPassword),
	}, nil
}

func ValidatePassword(plainPassword string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainPassword))
	return err == nil
}

func CreateToken(user User) string {
	now := time.Now()

	// valid for one day
	validUntil := now.Add(time.Hour * 24).Unix()

	claims := jwt.MapClaims{
		"user":    user.Username,
		"expires": validUntil,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims, nil)

	// do not do this, buddy
	// use some secret manager or sumn
	secret := "ULTRA MEGA SECRET STRING"

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return ""
	}

	return tokenString

}
