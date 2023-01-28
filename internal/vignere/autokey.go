package vignere

import (
	"github.com/ravielze/Crypto-1/internal"
	"github.com/ravielze/Crypto-1/internal/utils"
)

type (
	autokeyVigenere struct {
		key string
	}
)

func NewAutoKey(key string) internal.AlphabetCipher {
	return autokeyVigenere{
		key: key,
	}
}

func (a autokeyVigenere) Key() string {
	return a.key
}

func (a autokeyVigenere) Encrypt(plaintext string) string {
	key := utils.Equalize(plaintext, a.key)
	return NewStandard(key).Encrypt(plaintext)
}

func (a autokeyVigenere) Decrypt(ciphertext string) string {
	plaintext := ""
	j := 0
	key := a.key
	for _, c := range ciphertext {
		minus := utils.RuneBase(c)
		if minus == -1 {
			continue
		}

		keyChar := int32(key[j])
		keyMinus := utils.RuneBase(keyChar)
		c = (c-minus+keyMinus-keyChar+26)%26 + minus
		key += string(c)
		j = (j + 1) % len(key)
		plaintext += string(c)
	}
	return plaintext
}
