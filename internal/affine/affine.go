package affine

import (
	"fmt"
	"github.com/ravielze/Crypto-1/internal"
	"github.com/ravielze/Crypto-1/internal/utils"
	"strconv"
)

type (
	affine struct {
		m int
		n int
	}
)

func (a *affine) Metadata() map[string]any {
	return map[string]any{
		"m": a.m,
		"n": a.n,
	}
}

func New(m, n int) internal.BytesCipher {
	return &affine{m, n}
}

func (a *affine) Encrypt(plain []byte) []byte {
	if utils.GCD(a.m, 256) != 1 {
		fmt.Println("m and 256 are not coprime.")
		return nil
	}
	cipher := make([]byte, len(plain))
	for i, c := range plain {
		cipher[i] = byte((a.m*int(c) + a.n) % 256)
	}
	return cipher
}

func (a *affine) Decrypt(cipher []byte) []byte {
	if utils.GCD(a.m, 256) != 1 {
		fmt.Println("m and 256 are not coprime.")
		return nil
	}
	aInverse := 0
	for i := 0; i < 256; i++ {
		if (a.m*i)%256 == 1 {
			aInverse = i
			break
		}
	}
	plain := make([]byte, len(cipher))
	for i, c := range cipher {
		plain[i] = byte((aInverse * (int(c) + 256 - a.n)) % 256)
	}
	return plain
}

func (a *affine) Key() []byte {
	result := strconv.Itoa(a.m)
	result += ","
	result += strconv.Itoa(a.n)
	return []byte(result)
}
