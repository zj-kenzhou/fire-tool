package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func EncodePassword(password string) string {
	res, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(res)
}

func ValidatePassword(password, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(plainPassword))
	if err != nil {
		return false
	}
	return true
}
