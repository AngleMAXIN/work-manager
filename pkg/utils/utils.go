package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// MakePassword 生成密码
func MakePassword(rowPasswd string) (passWord string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(rowPasswd), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	passWord = string(hash)
	return
}

// CheckPassword 检查密码的正确性
func CheckPassword(inputPassword, savePassword string) (ok bool) {
	err := bcrypt.CompareHashAndPassword([]byte(savePassword), []byte(inputPassword))
	if err != nil {
		return
	}
	ok = true
	return
}

// func SendRespons(c)