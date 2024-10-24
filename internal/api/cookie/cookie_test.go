package cookie

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	crypt "github.com/rodeorm/shortener/internal/crypt"
)

func TestGetUserKeyFromCookie(t *testing.T) {
	validKey := strconv.Itoa(100)
	invalidKey := "invalid"
	validToken, _ := crypt.Encrypt(validKey)
	invalidToken, _ := crypt.Encrypt(invalidKey)

	tests := []struct {
		cookie *http.Cookie

		name string
		Key  string

		err bool
	}{
		{name: "обработка валидных куки", Key: validKey, cookie: &http.Cookie{
			Name:   "token",
			Value:  validToken,
			MaxAge: 10000,
		}, err: false},
		{name: "обработка невалидных куки", Key: invalidKey, cookie: &http.Cookie{
			Name:   "token",
			Value:  invalidToken,
			MaxAge: 10000,
		}, err: true},
		{name: "обработка пустых куки", cookie: &http.Cookie{}, err: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := http.Request{Header: map[string][]string{
				"Accept-Encoding": {"gzip, deflate"},
				"Accept-Language": {"en-us"},
			}}
			req.AddCookie(tt.cookie)

			key, err := GetUserKeyFromCookie(&req)

			if tt.err {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.Key, key)
		})
	}
}
