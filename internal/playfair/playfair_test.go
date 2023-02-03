package playfair

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlayfair_Encrypt(t *testing.T) {
	type args struct {
		key   string
		plain string
	}
	tests := []struct {
		tcNumber int
		args     args
		want     string
	}{
		{
			tcNumber: 1,
			args: args{
				plain: "KRIPTOGRAFIKLASIKDENGANCIPHERPLAYFAIR",
				key:   "LAMPION",
			},
			want: "FULIQCFSNRDUAMUMUKFOFMBDLIKFTAAMWHMLSW",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testcase Number #%d", tt.tcNumber), func(t *testing.T) {
			engine := NewPlayfair(tt.args.key)
			result := engine.Encrypt(tt.args.plain)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestPlayfair_Decrypt(t *testing.T) {
	type args struct {
		key    string
		cipher string
	}
	tests := []struct {
		tcNumber int
		args     args
		want     string
	}{
		{
			tcNumber: 1,
			args: args{
				cipher: "FULIQCFSNRDUAMUMUKFOFMBDLIKFTAAMWHMLSW",
				key:    "LAMPION",
			},
			want: "KRIPTOGRAFIKLASIKDENGANCIPHERPLAYFAIR",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testcase Number #%d", tt.tcNumber), func(t *testing.T) {
			engine := NewPlayfair(tt.args.key)
			result := engine.Decrypt(tt.args.cipher)
			assert.Equal(t, tt.want, result)
		})
	}
}
