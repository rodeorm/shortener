package repo

import (
	"fmt"

	"github.com/rodeorm/shortener/internal/core"
)

type memoryStorage struct {
	originalToShort map[string]string
	shortToOriginal map[string]string
	users           map[int]*core.User
	userURLPairs    map[int]*[]core.UserURLPair
}

// InsertShortURL принимает оригинальный URL, генерирует для него ключ и сохраняет соответствие оригинального URL и ключа (либо возвращает ранее созданный ключ)
func (s memoryStorage) InsertURL(URL, baseURL string, user *core.User) (*core.URL, error) {
	if !core.CheckURLValidity(URL) {
		return nil, fmt.Errorf("невалидный URL: %s", URL)
	}
	key, isExist := s.originalToShort[URL]
	if isExist {
		s.insertUserURLPair(baseURL+"/"+key, URL, user)
		return &core.URL{Key: key, HasBeenShorted: isExist}, nil
	}
	key, err := core.ReturnShortKey(5)
	if err != nil {
		return nil, err
	}

	s.originalToShort[URL] = key
	s.shortToOriginal[key] = URL

	s.insertUserURLPair(baseURL+"/"+key, URL, user)

	return &core.URL{Key: key, HasBeenShorted: false}, nil
}

// SelectOriginalURL принимает на вход короткий URL (относительный, без имени домена), извлекает из него ключ и возвращает оригинальный URL из хранилища
func (s memoryStorage) SelectOriginalURL(shortURL string) (*core.URL, error) {
	originalURL, isExist := s.shortToOriginal[shortURL]
	return &core.URL{Key: shortURL, HasBeenShorted: isExist, OriginalURL: originalURL}, nil
}

// InsertUser сохраняет нового пользователя или возвращает уже имеющегося в наличии
func (s memoryStorage) InsertUser(Key int) (*core.User, bool, error) {
	if Key == 0 {
		user := &core.User{Key: s.getNextFreeKey()}
		s.users[user.Key] = user
		return user, true, nil
	}
	user, isExist := s.users[Key]
	if !isExist {
		user = &core.User{Key: Key}
		s.users[Key] = user
		return user, true, nil
	}
	return user, false, nil
}

// InsertUserURLPair cохраняет информацию о том, что пользователь сокращал URL, если такой информации ранее не было
func (s memoryStorage) insertUserURLPair(shorten, origin string, user *core.User) error {

	URLPair := &core.UserURLPair{UserKey: user.Key, Short: shorten, Origin: origin}

	userURLPairs, isExist := s.userURLPairs[URLPair.UserKey]
	if !isExist {
		userURLPair := URLPair
		new := make([]core.UserURLPair, 0, 10)
		new = append(new, *userURLPair)
		s.userURLPairs[URLPair.UserKey] = &new
		return nil
	}

	for _, value := range *userURLPairs {
		if value.Origin == URLPair.Origin {
			return nil
		}
	}
	*s.userURLPairs[URLPair.UserKey] = append(*s.userURLPairs[URLPair.UserKey], *URLPair)

	return nil
}

func (s memoryStorage) SelectUserByKey(Key int) (*core.User, error) {
	user, isExist := s.users[Key]
	if !isExist {
		return nil, fmt.Errorf("нет пользователя с ключом: %d", Key)
	}
	return user, nil
}

// SelectUserURL возвращает перечень соответствий между оригинальным и коротким адресом для конкретного пользователя
func (s memoryStorage) SelectUserURLHistory(user *core.User) ([]core.UserURLPair, error) {
	if s.userURLPairs[user.Key] == nil {
		return nil, fmt.Errorf("нет истории")
	}
	return *s.userURLPairs[user.Key], nil
}

// getNextFreeKey возвращает ближайший свободный идентификатор пользователя
func (s memoryStorage) getNextFreeKey() int {
	var maxNumber int
	for maxNumber = range s.users {
		break
	}
	for n := range s.users {
		if n > maxNumber {
			maxNumber = n
		}
	}
	return maxNumber + 1
}

func (s memoryStorage) CloseConnection() {
	fmt.Println("Закрыто")
}

func (s memoryStorage) DeleteURLs(URLs []core.URL) error {
	return nil
}

func (s memoryStorage) PingDB() error {
	return nil
}
