package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetURLsFromString(t *testing.T) {

	type testInput struct {
		urls string
		user User
	}

	tests := []struct {
		name    string
		value   testInput
		want    int
		wantErr bool
	}{

		{
			name:    "Контроль ошибок. Пустая строка",
			value:   testInput{urls: "", user: User{Key: 1}},
			wantErr: true,
		},
		{
			name:    "Контроль ошибок. Нет пользователя",
			value:   testInput{urls: ""},
			wantErr: true,
		},
		{
			name:    "Успешный сценарий: 1 урл",
			value:   testInput{urls: "[\"6qxTVvsy\"]", user: User{Key: 1}},
			want:    1,
			wantErr: false,
		}, {
			name:    "Успешный сценарий: 3 урл",
			value:   testInput{urls: "[\"6qxTVvsy\", \"RTfd56hn\", \"Jlfd67ds\"]", user: User{Key: 1}},
			want:    3,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetURLsFromString(tt.value.urls, &tt.value.user)
			if !tt.wantErr {
				require.NoError(t, err)
				assert.Equal(t, tt.want, len(got))
				return
			}
			assert.Error(t, err)
		})
	}
}

func BenchmarkGetURLsFromString(b *testing.B) {
	type testInput struct {
		urls string
		user User
	}

	tests := []testInput{
		{urls: "[\"6qxTVvsy\", \"RTfd56hn\", \"Jlfd67ds\"]", user: User{Key: 1}},
		{urls: "[\"6qxTVvsy\", \"RTfd56hn\", \"Jlfd67ds\", \"Jlfd67ds\", \"Jlfd67ds\", \"Jlfd67ds\", \"Jlfd67ds\"]", user: User{Key: 2}},
	}

	for i := 0; i < 1000000; i++ {
		for _, v := range tests {
			GetURLsFromString(v.urls, &v.user)
		}
	}
}

func BenchmarkValidateURL(b *testing.B) {
	urls := []string{"http://www.yandex.ru", "https://www.ya.ru", "https://www.ya.com", "incorret.", "incorrect2"}

	b.Run("standart", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, v := range urls {
				CheckURLValidity(v)
			}
		}
	})
	b.Run("regexp", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, v := range urls {
				CheckURLValidityByRegExp(v)
			}
		}
	})

}
