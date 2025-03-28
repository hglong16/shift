package shift

import (
	"bytes"
	"errors"
)

func Encipher(plaintext []byte, key []byte) (ciphertext []byte) {
	ciphertext = make([]byte, len(plaintext))

	for i, b := range plaintext {
		ciphertext[i] = b + key[i%len(key)]
	}

	return
}

func Decipher(ciphertext []byte, key []byte) (plaintext []byte) {
	plaintext = make([]byte, len(ciphertext))
	for i, b := range ciphertext {
		plaintext[i] = b - key[i%len(key)]
	}

	return plaintext
}

const MaxKeyLen = 32

func Crack(ciphertext, crib []byte) (key []byte, err error) {
	for k := range min(MaxKeyLen, len(ciphertext)) {
		for guess := range 256 {
			result := ciphertext[k] - byte(guess)
			if result == crib[k] {
				key = append(key, byte(guess))
				break
			}
		}
		if bytes.Equal(crib, Decipher(ciphertext[:len(crib)], key)) {
			return key, nil
		}
	}
	return nil, errors.New("no key found")
}
