package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func MakeBcryptHashFromString(str string) ([]byte, error) {
	rawString := []byte(str)
	return bcrypt.GenerateFromPassword(rawString, 10)
}
