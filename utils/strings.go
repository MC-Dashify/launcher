package utils

import (
	"math/rand"
	"strings"

	"github.com/MC-Dashify/launcher/utils/logger"
	"golang.org/x/crypto/bcrypt"
)

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str || strings.HasPrefix(v, str) {
			return true
		}
	}
	return false
}

func GenerateRandomString(length int) string {
	charPool := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()/`~-_+=")

	b := make([]rune, length)
	for i := range b {
		b[i] = charPool[rand.Intn(len(charPool))]
	}

	return string(b)
}

func GenerateBCryptString(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error(err.Error())
	}
	return string(hash)
}