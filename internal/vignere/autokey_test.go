package vignere

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAutoKeyEncrypt(t *testing.T) {
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
				plain: "NEGARAPENGHASILMINYAK",
				key:   "INDO",
			},
			want: "VRJOEEVEEGWEFOSMAVJMS",
		},
		{
			tcNumber: 2,
			args: args{
				plain: "asisten baik dan ganteng #nilaibagusyak wkwkw",
				key:   "HALAL",
			},
			want: "hstseefjsboqbnokqtrtgabpnooiruazaqqcukg",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testcase Number #%d", tt.tcNumber), func(t *testing.T) {
			engine := NewAutoKey(tt.args.key)
			result := engine.Encrypt(tt.args.plain)
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestAutoKeyDecrypt(t *testing.T) {
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
				cipher: "VRJOEEVEEGWEFOSMAVJMS",
				key:    "INDO",
			},
			want: "NEGARAPENGHASILMINYAK",
		},
		{
			tcNumber: 2,
			args: args{
				cipher: "HSTSEEFJSBOQBNOKQTRTGABPNOOIRUAZAQQCUKG",
				key:    "HALAL",
			},
			want: "ASISTENBAIKDANGANTENGNILAIBAGUSYAKWKWKW",
		},
		{
			tcNumber: 3,
			args: args{
				cipher: "hstseefjsboqbnokqtrtgabpnooiruazaqqcukg",
				key:    "HALAL",
			}, want: "asistenbaikdangantengnilaibagusyakwkwkw",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Testcase Number #%d", tt.tcNumber), func(t *testing.T) {
			engine := NewAutoKey(tt.args.key)
			result := engine.Decrypt(tt.args.cipher)
			assert.Equal(t, tt.want, result)
		})
	}
}
