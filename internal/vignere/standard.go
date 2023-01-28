package vignere

import (
	"github.com/ravielze/Crypto-1/internal"
	"github.com/ravielze/Crypto-1/internal/utils"
)

type (
	standardVigenere struct {
		key string
	}
)

func (s standardVigenere) Metadata() map[string]any {
	return map[string]any{}
}

func (s standardVigenere) Key() string {
	return s.key
}

func NewStandard(key string) internal.AlphabetCipher {
	return standardVigenere{key}
}

func (s standardVigenere) Encrypt(plaintext string) string {
	ciphertext := ""
	key := utils.Normalize(s.key)
	j := 0
	for _, c := range plaintext {
		minus := utils.RuneBase(c)
		if minus == -1 {
			continue
		}

		keyChar := int32(key[j])
		keyMinus := utils.RuneBase(keyChar)
		c = (c-minus+keyChar-keyMinus)%26 + minus
		j = (j + 1) % len(key)
		ciphertext += string(c)
	}
	return ciphertext
}

func (s standardVigenere) Decrypt(ciphertext string) string {
	plaintext := ""
	j := 0
	key := s.key
	for _, c := range ciphertext {
		minus := utils.RuneBase(c)
		if minus == -1 {
			continue
		}

		keyChar := int32(key[j])
		keyMinus := utils.RuneBase(keyChar)
		c = (c-minus+keyMinus-keyChar+26)%26 + minus
		j = (j + 1) % len(key)
		plaintext += string(c)
	}
	return plaintext
}
