package crypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{
			name:  "проверка шифрования/дешифрования нормальной строки",
			value: "info",
		},
		{
			name:  "проверка шифрования/дешифрования пустой строки",
			value: "",
		},
		{
			name:  "проверка шифрования/дешифрования числа",
			value: "10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cypherValue, err := Encrypt(tt.value)
			require.NoError(t, err)
			decypherValue, err := Decrypt(cypherValue)
			require.NoError(t, err)
			assert.Equal(t, tt.value, decypherValue)
		})
	}
}
