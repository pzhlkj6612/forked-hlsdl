package hlsdl

import (
	"bytes"
	"testing"
)

func TestParseHexKey(t *testing.T) {
	expected := []byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcd, 0xef, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcd, 0xef}

	tests := []struct {
		name    string
		hexKey  string
		want    []byte
		wantErr bool
	}{
		{
			name:   "plain hex string",
			hexKey: "1234567890abcdef1234567890abcdef",
			want:   expected,
		},
		{
			name:   "hex string with 0x prefix",
			hexKey: "0x1234567890abcdef1234567890abcdef",
			want:   expected,
		},
		{
			name:   "hex string with 0X prefix",
			hexKey: "0X1234567890abcdef1234567890abcdef",
			want:   expected,
		},
		{
			name:   "uppercase hex",
			hexKey: "1234567890ABCDEF1234567890ABCDEF",
			want:   expected,
		},
		{
			name:    "invalid hex string",
			hexKey:  "not-a-hex-key",
			wantErr: true,
		},
		{
			name:    "empty string",
			hexKey:  "",
			want:    []byte{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseHexKey(tt.hexKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseHexKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !bytes.Equal(got, tt.want) {
				t.Errorf("parseHexKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
