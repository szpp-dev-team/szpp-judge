package user

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) []byte {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hashed
}

func VerifyPassword(hashed []byte, password string) error {
	return bcrypt.CompareHashAndPassword(hashed, []byte(password))
}
