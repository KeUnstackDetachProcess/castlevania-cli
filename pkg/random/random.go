package random

import (
	"crypto/rand"
)

// Generates a byte array which length depends on the "len" parameter
func GenerateByteArray(len int) ([]byte, error) {
	ret := make([]byte, len) // Our return value

	_, err := rand.Read(ret) // https://pkg.go.dev/crypto/rand#Read

	if err != nil {

		return nil, err
	}

	return ret, nil
}
