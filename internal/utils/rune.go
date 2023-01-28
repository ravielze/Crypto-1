package utils

import (
	"strings"
)

func RuneBase(char rune) int32 {
	representation := int(char)
	if representation >= 65 && representation <= 90 {
		return 65
	} else if representation >= 97 && representation <= 122 {
		return 97
	} else {
		return -1
	}
}

func Normalize(s string) string {
	var normalizedChar []rune
	for _, char := range s {
		charBase := RuneBase(char)
		if charBase != -1 {
			normalizedChar = append(normalizedChar, char)
		}
	}
	return strings.ToUpper(string(normalizedChar))
}

func Equalize(plaintext string, key string) string {
	result := Normalize(key)
	plaintext = Normalize(plaintext)
	i := 0
	for len(result) < len(plaintext) {
		result += string(plaintext[i])
		i = (i + 1) % len(plaintext)
	}
	return result
}
