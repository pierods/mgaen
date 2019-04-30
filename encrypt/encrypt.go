package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"
)

/*Seal encrypts and seals a []byte with aes256/gcm*/
func Seal(clearData, password []byte) ([]byte, error) {

	if len(clearData) == 0 {
		return []byte{}, errors.New("no input data, nothing to do")
	}

	if len(password) == 0 {
		return []byte{}, errors.New("empty password")
	}
	s256 := sha256.Sum256(password)
	block, err := aes.NewCipher(s256[:])
	if err != nil {
		return []byte{}, err
	}
	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, err
	}

	nonce := make([]byte, aesGcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return []byte{}, err
	}
	cipherText := aesGcm.Seal(nil, nonce, clearData, nil)
	// https://crypto.stackexchange.com/questions/57895/would-it-be-safe-to-store-gcm-nonce-in-the-encrypted-output
	nonceAndEncBytes := append(nonce, cipherText...)
	return nonceAndEncBytes, nil
}
