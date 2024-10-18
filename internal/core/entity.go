package core

/****
 For minimizing the amount of padding bytes, you must lay out the fields from highest allocation to lowest allocation.
 This will push any necessary padding bytes down to the bottom of the struct and reduce the total number of padding bytes necessary
***/

// URL - структура для данных URL
type URL struct {
	OriginalURL    string // Оригинальный урл
	Key            string `json:"url,omitempty" db:"key"` // Ключ, использованный при сокращении
	UserKey        int    `db:"user_key"`                 // Пользователь, который сократил URL
	HasBeenShorted bool   // Признак, что сокращали ранее
	HasBeenDeleted bool   // Признал, что был удален
}

// ShortenURL - структура для сериализации сокращенного УРЛ
type ShortenURL struct {
	Key string `json:"result,omitempty"` // Ключ, использованный при сокращении
}

// URLPair - структура для серилизации cоответствие пар URL в end-point api shorten
type URLPair struct {
	Origin string `json:"origin,omitempty"` // Оригинальный URL
	Short  string `json:"short,omitempty"`  // Сокращенный URL
}

// UserURLPair - структура для сериализации соответствия пар URL в end-point api user url
type UserURLPair struct {
	Short   string `json:"short_url,omitempty"`    // Сокращенный URL
	Origin  string `json:"original_url,omitempty"` // Оригинальный URL
	UserKey int    `json:"-"`                      // Уникальный идентификатор пользователя
}

// UrlWithCorrelationRequest - структура для сериализации URL в end-point api/shorten/batch
type URLWithCorrelationRequest struct {
	CorID  string `json:"correlation_id,omitempty"` // Идентификатор сокращенного URL
	Origin string `json:"original_url,omitempty"`   // Оригинальный URL
}

// UrlWithCorrelationResponse - структура для сериализации  множество URL в end-point api/shorten/batch
type URLWithCorrelationResponse struct {
	CorID string `json:"correlation_id,omitempty"` // Идентификатор сокращенного URL
	Short string `json:"short_url,omitempty"`      // Оригинальный URL
}

// User - пользователь сервиса
type User struct {
	Urls           []UserURLPair // Сокращенные пользователем URL
	Key            int           // Уникальный идентификатор пользователя
	WasUnathorized bool          // Признак того, что пользователь был создан автоматически, после того как не получилось авторизовать его через куки
}
