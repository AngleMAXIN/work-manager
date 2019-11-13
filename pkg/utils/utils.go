package utils

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
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
var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}
