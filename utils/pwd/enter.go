package pwd

import (
	"ys_go/global"

	"golang.org/x/crypto/bcrypt"
)

// 加密密码
func GenerateFromPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		global.Log.Error("GenerateFromPassword", err)
		return "", err
	}
	return string(hashedPassword), nil
}

// 校验密码
func CompareHashAndPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
