package internal

type (
	AlphabetCipher interface {
		Encrypt(plaintext string) string
		Decrypt(ciphertext string) string
		Key() string
		Metadata() map[string]any
	}

	BytesCipher interface {
		Encrypt(plain []byte) []byte
		Decrypt(cipher []byte) []byte
		Key() []byte
		Metadata() map[string]any
	}
)
