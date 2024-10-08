package core

type URL struct {
	Key            string `json:"url,omitempty" db:"key"` //Ключ, использованный при сокращении
	UserKey        int    `db:"user_key"`                 // Пользователь, который сократил URL
	OriginalURL    string // Оригинальный урл
	HasBeenShorted bool   // Признак, что сокращали ранее
	HasBeenDeleted bool   // Признал, что был удален
}

type ShortenURL struct {
	Key string `json:"result,omitempty"`
}

// URLPair - cоответствие пар URL для api shorten
type URLPair struct {
	Origin string `json:"origin,omitempty"`
	Short  string `json:"short,omitempty"`
}

// UserURLPair - соответствие пар URL для api user url
type UserURLPair struct {
	UserKey int    `json:"-"`
	Short   string `json:"short_url,omitempty"`
	Origin  string `json:"original_url,omitempty"`
	// isDeleted bool   `json:"-"`
}

// UrlWithCorrelationRequest множество url для запроса api/shorten/batch
type URLWithCorrelationRequest struct {
	CorID  string `json:"correlation_id,omitempty"`
	Origin string `json:"original_url,omitempty"`
}

// UrlWithCorrelationResponse множество url для ответа api/shorten/batch
type URLWithCorrelationResponse struct {
	CorID string `json:"correlation_id,omitempty"`
	Short string `json:"short_url,omitempty"`
}

type User struct {
	Key            int
	Urls           []UserURLPair
	WasUnathorized bool
}
