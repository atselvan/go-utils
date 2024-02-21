package security

import (
	"encoding/base64"

	"github.com/atselvan/go-utils/utils/errors"
)

// Base64Encode base64 encodes a string and returns it.
func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Base64Decode base64 decodes a string and returns it.
func Base64Decode(s string) (string, *errors.Error) {
	passwordBytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", errors.New(errors.ErrCodeBase64DecodeError, 0, err.Error())
	}
	return string(passwordBytes), nil
}
