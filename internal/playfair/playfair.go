package playfair

import (
	"github.com/ravielze/Crypto-1/internal"
	"github.com/ravielze/Crypto-1/internal/utils"
	"strings"
)

type (
	playfair struct {
		key string
	}
)

func NewPlayfair(key string) internal.AlphabetCipher {
	return &playfair{key}
}

func (p *playfair) Metadata() map[string]any {
	return map[string]any{}
}

func (p *playfair) Key() string {
	return p.key
}

func (p *playfair) Encrypt(plaintext string) string {
	key := utils.Normalize(p.key)
	plaintext = utils.Normalize(plaintext)
	table := utils.GenerateTable(key)
	offset := 1

	var result []rune
	for i := 0; i < len(plaintext); {
		var char1, char2 rune
		if i == len(plaintext)-1 {
			char1, char2 = rune(plaintext[i]), 'X'
		} else {
			char1, char2 = rune(plaintext[i]), rune(plaintext[i+1])
		}

		if char1 == char2 {
			char2 = 'X'
			i++
		} else {
			i += 2
		}

		i1, i2 := table.GetIndex(char1, char2)
		if table.IsSameRow(i1, i2) {
			i1, i2 = table.ShiftHorizontal(i1, i2, offset)
		} else if table.IsSameColumn(i1, i2) {
			i1, i2 = table.ShiftVertical(i1, i2, offset)
		} else {
			i1, i2 = table.ShiftRectangle(i1, i2)
		}

		result = append(result, table[i1], table[i2])
	}
	return string(result)
}

func (p *playfair) Decrypt(ciphertext string) string {
	key := utils.Normalize(p.key)
	ciphertext = utils.Normalize(ciphertext)
	table := utils.GenerateTable(key)
	offset := 1

	var result []rune
	for i := 0; i < len(ciphertext); i += 2 {
		char1, char2 := rune(ciphertext[i]), rune(ciphertext[i+1])

		i1, i2 := table.GetIndex(char1, char2)
		if table.IsSameRow(i1, i2) {
			i1, i2 = table.ShiftHorizontal(i1, i2, -1*offset)
		} else if table.IsSameColumn(i1, i2) {
			i1, i2 = table.ShiftVertical(i1, i2, -1*offset)
		} else {
			i1, i2 = table.ShiftRectangle(i1, i2)
		}

		result = append(result, table[i1], table[i2])
	}
	return strings.ReplaceAll(string(result), "X", "")
}
