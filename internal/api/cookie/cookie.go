package cookie

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/rodeorm/shortener/internal/crypt"
)

// GetUserKeyFromCookie получает идентфикатор пользователя из куки "токен"
func GetUserKeyFromCookie(r *http.Request) (string, error) {
	tokenCookie, err := r.Cookie("token")
	if err != nil {
		return "", err
	}
	if tokenCookie.Value == "" {
		return "", fmt.Errorf("не найдено актуальных cookie")
	}
	userKey, err := crypt.Decrypt(tokenCookie.Value)
	if err != nil {
		return "", err
	}

	_, err = strconv.Atoi(userKey)

	if err != nil {
		return "", err
	}

	return userKey, nil
}

// PutUserKeyToCookie помещает идентификатор пользователя в  куки "токен"
func PutUserKeyToCookie(Key string) *http.Cookie {
	val, err := crypt.Encrypt(Key)
	if err != nil {
		panic(err)
	}

	cookie := &http.Cookie{
		Name:   "token",
		Value:  val,
		MaxAge: 10000,
	}
	return cookie
}
