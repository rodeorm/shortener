package repo

import (
	"fmt"
	"strconv"

	"github.com/rodeorm/shortener/internal/logic"
)

type memoryStorage struct {
	originalToShort map[string]string
	shortToOriginal map[string]string
	users           map[int]*User
	userURLPairs    map[int]*[]UserURLPair
}

//InsertShortURL принимает оригинальный URL, генерирует для него ключ и сохраняет соответствие оригинального URL и ключа (либо возвращает ранее созданный ключ)
func (s memoryStorage) InsertURL(URL, baseURL, userKey string) (string, bool, error) {
	if !logic.CheckURLValidity(URL) {
		return "", false, fmt.Errorf("невалидный URL: %s", URL)
	}
	key, isExist := s.originalToShort[URL]
	if isExist {
		s.insertUserURLPair(userKey, baseURL+"/"+key, URL)
		return key, true, nil
	}
	key, _ = logic.ReturnShortKey(5)

	s.originalToShort[URL] = key
	s.shortToOriginal[key] = URL

	s.insertUserURLPair(userKey, baseURL+"/"+key, URL)

	return key, false, nil
}

//SelectOriginalURL принимает на вход короткий URL (относительный, без имени домена), извлекает из него ключ и возвращает оригинальный URL из хранилища
func (s memoryStorage) SelectOriginalURL(shortURL string) (string, bool, bool, error) {
	originalURL, isExist := s.shortToOriginal[shortURL]
	return originalURL, isExist, false, nil
}

//InsertUser сохраняет нового пользователя или возвращает уже имеющегося в наличии
func (s memoryStorage) InsertUser(Key int) (*User, error) {
	if Key == 0 {
		user := &User{Key: s.getNextFreeKey()}
		s.users[user.Key] = user
		return user, nil
	}
	user, isExist := s.users[Key]
	if !isExist {
		user = &User{Key: Key}
		s.users[Key] = user
	}
	return user, nil
}

//InsertUserURLPair cохраняет информацию о том, что пользователь сокращал URL, если такой информации ранее не было
func (s memoryStorage) insertUserURLPair(userKey, shorten, origin string) error {
	userID, err := strconv.Atoi(userKey)
	if err != nil {
		return fmt.Errorf("ошибка обработки идентификатора пользователя: %s", err)
	}

	URLPair := &UserURLPair{UserKey: userID, Short: shorten, Origin: origin}

	userURLPairs, isExist := s.userURLPairs[URLPair.UserKey]
	if !isExist {
		userURLPair := *URLPair
		new := make([]UserURLPair, 0, 10)
		new = append(new, userURLPair)
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

func (s memoryStorage) SelectUserByKey(Key int) (*User, error) {
	user, isExist := s.users[Key]
	if !isExist {
		return nil, fmt.Errorf("нет пользователя с ключом: %d", Key)
	}
	return user, nil
}

//SelectUserURL возвращает перечень соответствий между оригинальным и коротким адресом для конкретного пользователя
func (s memoryStorage) SelectUserURLHistory(Key int) (*[]UserURLPair, error) {
	if s.userURLPairs[Key] == nil {
		return nil, fmt.Errorf("нет истории")
	}
	return s.userURLPairs[Key], nil
}

//getNextFreeKey возвращает ближайший свободный идентификатор пользователя
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

func (s memoryStorage) DeleteURLs(URL, userKey string) (bool, error) {
	return true, nil
}
