package controller

import (
	"crypto/rand"
	"encoding/base64"
)

// createRandomBookID creates URL safe random book ids
func createRandomBookID() (string, error) {

	randomBytes := make([]byte, 32)
	// TODO:rand.Read might be a single shared resource
	// think about it later.
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	bookid := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(randomBytes)

	return bookid, nil
}
