package vignere

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStandardEncrypt(t *testing.T) {
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
				plain: "thisplaintext",
				key:   "SONY",
			},
			want: "lvvqhzngfhrvl",
		},
		{
			tcNumber: 2,
			args: args{
				plain: "HAIHAIHALO SEMUA",
				key:   "BEBEK",
			},
			want: "IEJLKJLBPYTINYK",
		},
		{
			tcNumber: 3,
			args: args{
				plain: "HAIhaiHALO SEMUA",
				key:   "BEBEK",
			},
			want: "IEJlkjLBPYTINYK",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testcase Number #%d", tt.tcNumber), func(t *testing.T) {
			engine := NewStandard(tt.args.key)
			result := engine.Encrypt(tt.args.plain)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestStandardDecrypt(t *testing.T) {
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
				cipher: "lvvqhzngfhrvl",
				key:    "SONY",
			},
			want: "thisplaintext",
		},
		{
			tcNumber: 2,
			args: args{
				cipher: "IEJLKJLBPYTINYK",
				key:    "BEBEK",
			},
			want: "HAIHAIHALOSEMUA",
		},
		{
			tcNumber: 3,
			args: args{
				cipher: "IEJLKJLBPYtinyk",
				key:    "BEBEK",
			},
			want: "HAIHAIHALOsemua",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testcase Number #%d", tt.tcNumber), func(t *testing.T) {
			engine := NewStandard(tt.args.key)
			result := engine.Decrypt(tt.args.cipher)
			assert.Equal(t, tt.want, result)
		})
	}
}
