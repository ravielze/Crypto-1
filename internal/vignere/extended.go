package vignere

import (
	"github.com/ravielze/Crypto-1/internal"
	"github.com/ravielze/Crypto-1/internal/utils"
)

type (
	extended struct {
		key []byte
	}

	extendedStr struct {
		operator internal.BytesCipher
	}
)

func (e *extendedStr) Encrypt(plaintext string) string {
	return string(e.operator.Encrypt([]byte(plaintext)))
}

func (e *extendedStr) Decrypt(ciphertext string) string {
	return string(e.operator.Decrypt([]byte(ciphertext)))
}

func (e *extendedStr) Key() string {
	return string(e.operator.Key())
}

func (e *extendedStr) Metadata() map[string]any {
	return e.operator.Metadata()
}

func (e *extended) Encrypt(plain []byte) []byte {
	var result []byte
	key := e.key
	for i, char := range plain {
		mod := utils.Modulus(int(char+key[i%len(key)]), 256)
		result = append(result, byte(mod))
	}
	return result
}

func (e *extended) Decrypt(cipher []byte) []byte {
	var result []byte
	key := e.key
	for i, char := range cipher {
		mod := utils.Modulus(int(char-key[i%len(key)]), 256)
		result = append(result, byte(mod))
	}

	return result
}

func (e *extended) Key() []byte {
	return e.key
}

func (e *extended) Metadata() map[string]any {
	return map[string]any{}
}

func NewExtended(key []byte) internal.BytesCipher {
	return &extended{key}
}

func NewExtendedString(key string) internal.AlphabetCipher {
	return &extendedStr{operator: NewExtended([]byte(key))}
}
