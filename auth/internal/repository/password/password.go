package password

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword used to hash password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// ComparePassword return an error or nil after comparing
func ComparePassword(providedPassword string, userPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true

	if err != nil {
		check = false
	}
	return check
}
