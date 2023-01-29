package vignere

import (
	"github.com/ravielze/Crypto-1/internal"
	"github.com/ravielze/Crypto-1/internal/utils"
)

type (
	autokey struct {
		key         string
		lastAutokey string
	}
)

func (a *autokey) Metadata() map[string]any {
	return map[string]any{
		"last_autokey": a.lastAutokey,
	}
}

func NewAutoKey(key string) internal.AlphabetCipher {
	return &autokey{
		key: key,
	}
}

func (a *autokey) Key() string {
	return a.key
}

func (a *autokey) Encrypt(plaintext string) string {
	key := utils.Equalize(plaintext, a.key)
	a.lastAutokey = key
	return NewStandard(key).Encrypt(plaintext)
}

func (a *autokey) Decrypt(ciphertext string) string {
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
	a.lastAutokey = key
	return plaintext
}
