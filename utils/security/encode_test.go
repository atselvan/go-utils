package security

import (
	"testing"

	"github.com/atselvan/go-utils/utils/errors"
	"github.com/stretchr/testify/assert"
)

const (
	password       = "password"
	base64Password = "cGFzc3dvcmQ="
)

func TestBase64Encode(t *testing.T) {
	encryptedPassword := Base64Encode(password)
	assert.Equal(t, base64Password, encryptedPassword)
}

func TestBase64Decode(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		decryptedPassword, cErr := Base64Decode(base64Password)
		assert.Equal(t, password, decryptedPassword)
		assert.Nil(t, cErr)
	})

	t.Run("error", func(t *testing.T) {
		decryptedPassword, cErr := Base64Decode("cGFzc3dvcmQ")
		assert.Empty(t, decryptedPassword)
		assert.NotNil(t, cErr)
		assert.Equal(t, errors.ErrCodeBase64DecodeError, cErr.Code)
		assert.Equal(t, "illegal base64 data at input byte 8", cErr.Message)
	})
}
