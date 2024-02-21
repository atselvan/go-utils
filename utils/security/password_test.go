package security

import (
	"fmt"
	"testing"

	"github.com/atselvan/go-utils/utils/errors"
	"github.com/stretchr/testify/assert"
)

func TestGetRandomPassword(t *testing.T) {
	p0 := GetRandomPassword()
	assert.NotEmpty(t, p0)

	p1 := GetRandomPassword()
	p2 := GetRandomPassword()
	assert.NotEqual(t, p1, p2)
}

func TestVerifyPassword(t *testing.T) {
	// Test if password has a number in it
	t.Run("number in password", func(t *testing.T) {
		password := "@Password"
		cErr := VerifyPassword(password)
		assert.Equal(t, errors.ErrCodeInvalidPassword, cErr.Code)
		assert.Equal(t, errors.ErrMsg[errors.ErrCodeInvalidPassword], cErr.Message)
	})

	// Test if password has an upper case letter
	t.Run("upper case letter in password", func(t *testing.T) {
		password := "@password123"
		cErr := VerifyPassword(password)
		assert.Equal(t, errors.ErrCodeInvalidPassword, cErr.Code)
		assert.Equal(t, errors.ErrMsg[errors.ErrCodeInvalidPassword], cErr.Message)
	})

	// Test if password has a lower case letter
	t.Run("lower case letter in password", func(t *testing.T) {
		password := "@PASSWORD123"
		cErr := VerifyPassword(password)
		assert.Equal(t, errors.ErrCodeInvalidPassword, cErr.Code)
		assert.Equal(t, errors.ErrMsg[errors.ErrCodeInvalidPassword], cErr.Message)
	})

	// Test if password has a special letter
	t.Run("special case character in password", func(t *testing.T) {
		password := "PASSWORD123"
		cErr := VerifyPassword(password)
		assert.Equal(t, errors.ErrCodeInvalidPassword, cErr.Code)
		assert.Equal(t, errors.ErrMsg[errors.ErrCodeInvalidPassword], cErr.Message)
	})

	// Test password length
	t.Run("password length", func(t *testing.T) {
		password := "@123"
		cErr := VerifyPassword(password)
		assert.Equal(t, errors.ErrCodeInvalidPassword, cErr.Code)
		assert.Equal(t, errors.ErrMsg[errors.ErrCodeInvalidPassword], cErr.Message)
	})

	// Test valid password
	t.Run("valid password", func(t *testing.T) {
		password := "somePass@123"
		cErr := VerifyPassword(password)
		assert.Nil(t, cErr)
	})
}

func TestEncryptPassword(t *testing.T) {
	password := "somepassword@123"
	ePass, cErr := EncryptPassword(password, "")
	assert.Nil(t, cErr)
	assert.NotEqual(t, password, ePass)
}

func TestDecryptPassword(t *testing.T) {
	password := "somepassword@123"
	passphrase := "something"

	ePass, cErr := EncryptPassword(password, "")
	assert.Nil(t, cErr)
	dPass, cErr := DecryptPassword(ePass, "")
	assert.Nil(t, cErr)
	assert.Equal(t, password, dPass)

	ePass, cErr = EncryptPassword(password, passphrase)
	assert.Nil(t, cErr)
	dPass, cErr = DecryptPassword(ePass, passphrase)
	assert.Nil(t, cErr)
	assert.Equal(t, password, dPass)

	dPass, cErr = DecryptPassword("====ddwf=", "notValid")
	assert.NotNil(t, cErr)
	assert.Equal(t, errors.ErrCodePasswordDecryptionError, cErr.Code)
	assert.Contains(t, cErr.Message, fmt.Sprintf(errors.ErrMsg[errors.ErrCodePasswordDecryptionError], ""))

	dPass, cErr = DecryptPassword(ePass, "notValid")
	assert.NotNil(t, cErr)
	assert.Equal(t, errors.ErrCodePasswordDecryptionError, cErr.Code)
	assert.Contains(t, cErr.Message, fmt.Sprintf(errors.ErrMsg[errors.ErrCodePasswordDecryptionError], ""))
}
