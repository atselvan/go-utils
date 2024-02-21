package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
	mr "math/rand"
	"time"
	"unicode"

	"github.com/atselvan/go-utils/utils/errors"
)

// GetRandomPassword generates a random string of upper + lower case alphabets and digits
// which is 23 bits long and returns the string
func GetRandomPassword() string {
	mr.Seed(time.Now().UnixNano())
	digits := "0123456789"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" + digits
	length := 23
	buf := make([]byte, length)
	buf[0] = digits[mr.Intn(len(digits))]
	for i := 1; i < length; i++ {
		buf[i] = all[mr.Intn(len(all))]
	}
	mr.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	return string(buf)
}

// VerifyPassword checks if the password is strong.
// The function checks if the password string is at least 8 character long, has an uppercase character,
// a lower case character, a number and a special character.
// If one of this condition fails the function returns an invalid password error.
func VerifyPassword(password string) *errors.Error {
	var (
		numOfLetters                  = 0
		number, upper, lower, special bool
	)
	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			number = true
			numOfLetters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
			numOfLetters++
		case unicode.IsUpper(c):
			upper = true
			numOfLetters++
		case unicode.IsLower(c) || c == ' ':
			lower = true
			numOfLetters++
		}
	}
	if numOfLetters > 8 && number && upper && lower && special {
		return nil
	} else {
		return errors.New(
			errors.ErrCodeInvalidPassword,
			0,
			errors.ErrMsg[errors.ErrCodeInvalidPassword],
		)
	}
}

// createSHA256Hash creates a SHA256 checksum of the provided string.
func createSHA256Hash(key string) []byte {
	hash := sha256.Sum256([]byte(key))
	return hash[:]
}

// EncryptPassword encrypts the data provided.
// The encryption can also be further secured using a passphrase.
func EncryptPassword(data, passphrase string) (string, *errors.Error) {
	block, _ := aes.NewCipher(createSHA256Hash(passphrase))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.Newf(
			errors.ErrCodePasswordEncryptionError,
			0,
			errors.ErrMsg[errors.ErrCodePasswordEncryptionError],
			err.Error(),
		)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", errors.Newf(
			errors.ErrCodePasswordEncryptionError,
			0,
			errors.ErrMsg[errors.ErrCodePasswordEncryptionError],
			err.Error(),
		)
	}
	ciphertext := gcm.Seal(nonce, nonce, []byte(data), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptPassword decrypted the data provided.
// If a passphrase was used to encrypt the password then the same passphrase needs to be passes
// to decrypt the password correctly.
func DecryptPassword(data, passphrase string) (string, *errors.Error) {
	bData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", errors.Newf(
			errors.ErrCodePasswordDecryptionError,
			0,
			errors.ErrMsg[errors.ErrCodePasswordDecryptionError],
			err.Error(),
		)
	}
	key := createSHA256Hash(passphrase)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", errors.Newf(
			errors.ErrCodePasswordDecryptionError,
			0,
			errors.ErrMsg[errors.ErrCodePasswordDecryptionError],
			err.Error(),
		)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.Newf(
			errors.ErrCodePasswordDecryptionError,
			0,
			errors.ErrMsg[errors.ErrCodePasswordDecryptionError],
			err.Error(),
		)
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := bData[:nonceSize], bData[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", errors.Newf(
			errors.ErrCodePasswordDecryptionError,
			0,
			errors.ErrMsg[errors.ErrCodePasswordDecryptionError],
			err.Error(),
		)
	}
	return string(plaintext), nil
}
