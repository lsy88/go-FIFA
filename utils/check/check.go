package check

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func CheckName(v string) bool {
	if v == "" {
		return false
	}
	trim := strings.TrimSpace(v)
	if len(trim) <= 6 {
		return false
	}
	return true
}

func PassWordEncrypted(value string) string {
	passWord := []byte(string(value))
	hashedPassword, err := bcrypt.GenerateFromPassword(passWord, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s", hashedPassword)
	
}
