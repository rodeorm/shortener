package core

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

// GetClearURL делает URL строчным, убирает наименование домена
func GetClearURL(s string, d string) string {
	s = strings.ToLower(s)
	return strings.Replace(s, d, "", 1)
}

// CheckURLValidity проверяет URL на корректность
func CheckURLValidity(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}

// CheckURLValidityByRegExp проверяет URL на корректность через regexp
func CheckURLValidityByRegExp(u string) bool {
	var urlRegex = regexp.MustCompile(`^(http|https)://[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}(/[a-zA-Z0-9-._~:?#@!$&'()*+,;=]*)*$`)
	return urlRegex.MatchString(u)
}

// GetURLsFromString получает строку, содержащую URL и пользователя.
// Возвращает слайс URL или ошибку
func GetURLsFromString(s string, u *User) ([]URL, error) {
	if u.Key <= 0 {
		return nil, fmt.Errorf("некорректный пользователь: %d", u.Key)
	}
	if s == "" {
		return nil, fmt.Errorf("пустая строка с url")
	}

	var replacer = strings.NewReplacer(" ", "", "\"", "", "[", "", "]", "")
	urls := make([]URL, 0)
	for _, v := range strings.Split(replacer.Replace(s), ",") {
		urls = append(urls, URL{Key: v, UserKey: u.Key})
	}
	return urls, nil
}
