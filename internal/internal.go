package internal

type (
	AlphabetCipher interface {
		Encrypt(plaintext string) string
		Decrypt(ciphertext string) string
		Key() string
	}

	Cipher interface {
		Encrypt(plain []byte) []byte
		Decrypt(cipher []byte) []byte
	}
)
