package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"testing"
)

func TestSeal(t *testing.T) {

	_, err := Seal([]byte{}, []byte{1})
	if err == nil {
		t.Fatal("Should not accept an empty slice of data")
	}
	_, err = Seal([]byte{1}, []byte{})
	if err == nil {
		t.Fatal("Should not accept an empty password")
	}
	t.Log([]byte("abcdefghi"))
	enc, err := Seal([]byte("abcdefghi"), []byte("0123456789"))
	if err != nil {
		t.Fatal(err)
	}
	dec := decrypt(t, enc, []byte("0123456789"))
	if dec != "abcdefghi" {
		t.Fatal("Should be able to encode an input string")
	}
}

func decrypt(t *testing.T, data, password []byte) string {

	s256 := sha256.Sum256(password)
	block, err := aes.NewCipher(s256[:])
	if err != nil {
		t.Fatal(err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		t.Fatal(err)
	}

	nonce, cipherText := data[:aesGCM.NonceSize()], data[aesGCM.NonceSize():]

	decr, err := aesGCM.Open(nil, nonce, cipherText, nil)

	if err != nil {
		t.Fatal(err)
	}
	return string(decr)
}
