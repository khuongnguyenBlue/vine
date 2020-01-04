package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func EncryptPassword(input string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input), 4)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(hash), nil
}

func ComparePassword(encryptedPassword string, input string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(input))
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
