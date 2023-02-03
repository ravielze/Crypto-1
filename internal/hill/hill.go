package hill

import (
	"github.com/ravielze/Crypto-1/internal"
	"github.com/ravielze/Crypto-1/internal/utils"
	"math"
)

type (
	hill struct {
		key string
	}
)

func NewHill(key string) internal.AlphabetCipher {
	return &hill{key}
}

func (h *hill) Key() string {
	return h.key
}

func (h *hill) Metadata() map[string]any {
	return map[string]any{}
}

func (h *hill) Encrypt(plaintext string) string {
	key := utils.Normalize(h.key)
	plaintext = utils.Normalize(plaintext)

	if !utils.IsQuadratic(len(key)) {
		panic("Error")
	}
	size := int(math.Sqrt(float64(len(key))))

	keyMatrix := utils.GenerateKeyMatrix(key, size)

	segmentedPlaintext := utils.GenerateSegmentedText(plaintext, size)

	return Multiply(segmentedPlaintext, keyMatrix, size)
}

func (h *hill) Decrypt(ciphertext string) string {
	key := utils.Normalize(h.key)
	ciphertext = utils.Normalize(ciphertext)

	if !utils.IsQuadratic(len(key)) {
		panic("Error")
	}
	size := int(math.Sqrt(float64(len(key))))

	keyMatrix := utils.GenerateKeyMatrix(key, size)

	inverse := keyMatrix.Inverse()

	segmentedCipherText := utils.GenerateSegmentedText(ciphertext, size)

	return Multiply(segmentedCipherText, inverse, size)
}

func Multiply(segmentedText []utils.Matrix, matrix utils.Matrix, size int) string {
	var result []rune
	for _, segment := range segmentedText {
		tmp := matrix.Multiply(segment)
		for i := 0; i < size; i++ {
			char := utils.Modulo(tmp[i][0], 26) + 65
			result = append(result, rune(char))
		}
	}
	return string(result)
}
