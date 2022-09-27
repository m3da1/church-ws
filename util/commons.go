package util

import (
	b64 "encoding/base64"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

const DateOnlyformatter = "2006-01-02"

func GenerateResponse(code int, msg string, result interface{}) gin.H {
	return gin.H{
		"code":    code,
		"message": msg,
		"result":  result,
	}
}

func EncryptPassword(password string) string {
	return b64.StdEncoding.EncodeToString([]byte(password))
}

func VerifyPassword(currentPassword, hashedPassword string) (bool, error) {
	decryptedPassword, err := b64.StdEncoding.DecodeString(hashedPassword)
	if err != nil {
		return false, fmt.Errorf("password decryption error; %+v", err.Error())
	}
	if currentPassword != string(decryptedPassword) {
		return false, fmt.Errorf("invalid password")
	}
	return true, nil
}

func ParseDate(datestr string) (time.Time, error) {
	date, err := time.Parse(DateOnlyformatter, datestr)
	if err != nil {
		return time.Now(), err
	}
	return date, nil
}
