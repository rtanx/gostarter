package security

import (
	"github.com/rtanx/gostarter/internal/infrastructure/logger"
	"golang.org/x/crypto/bcrypt"
)

const pwdHashCost = 12

func HashPassword(pwd string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pwd), pwdHashCost)
	if err != nil {
		logger.Error("unexpected error while hashing password", logger.Err(err))
		return "", err
	}
	return string(hashed), nil
}

func CheckPassword(hashedPwd, pwd string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd)) == nil
}
