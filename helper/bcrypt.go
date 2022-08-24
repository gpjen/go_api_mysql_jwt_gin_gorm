package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func HasshAndSalt(pwd []byte) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	return string(hashed), err
}

func ComparePasword(pwd string, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	if err != nil {
		return false, err
	}
	return true, nil
}
