package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// In a real application, you would use a library like bcrypt to hash the password
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPasswordHash(password, hash string) bool {
	// In a real application, you would use a library like bcrypt to check the password hash
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
