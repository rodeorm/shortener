package repo

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/rodeorm/shortener/internal/core"
)

// InsertShortURL принимает оригинальный URL, генерирует для него ключ и сохраняет соответствие оригинального URL и ключа (либо возвращает ранее созданный ключ)
func (s fileStorage) InsertURL(URL, baseURL string, user *core.User) (*core.URL, error) {
	url := core.URL{OriginalURL: core.GetClearURL(URL, "")}

	if !core.CheckURLValidity(URL) {
		return nil, fmt.Errorf("невалидный URL: %s", URL)
	}

	key, isShorted, err := s.getShortlURLFromFile(URL)
	if err != nil {
		return nil, fmt.Errorf("ошибка при открытии файла в InsertURL %w", err)
	}
	url.Key = key
	url.HasBeenShorted = isShorted

	if url.HasBeenShorted {
		return &url, nil
	}
	url.Key, _ = core.ReturnShortKey(5)

	f, err := os.OpenFile(s.filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return nil, fmt.Errorf("ошибка при открытии файла в InsertURL %w", err)
	}
	defer f.Close()

	pair := core.URLPair{Origin: url.OriginalURL, Short: url.Key}
	data, err := json.Marshal(pair)
	if err != nil {
		return nil, fmt.Errorf("ошибка при открытии маршалининге урл %s в InsertURL %w", pair, err)
	}

	s.insertUserURLPair(baseURL+"/"+url.Key, URL, user)
	data = append(data, '\n')
	_, err = f.Write(data)

	return &url, err
}

// getShortlURLFromFile возвращает из файла сокращенный URL по оригинальному URL
func (s fileStorage) getShortlURLFromFile(URL string) (string, bool, error) {

	file, err := os.Open(s.filePath)
	if err != nil {
		return "", false, fmt.Errorf("ошибка при открытии файла в getShortlURLFromFile %w", err)
	}
	defer file.Close()

	var up core.URLPair
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		json.Unmarshal(scanner.Bytes(), &up)
		if up.Origin == URL {
			return up.Short, true, nil
		}
	}

	return "", false, nil
}

// SelectOriginalURL принимает на вход короткий URL (относительный, без имени домена), извлекает из него ключ и возвращает оригинальный URL из хранилища
func (s fileStorage) SelectOriginalURL(shortURL string) (*core.URL, error) {

	file, err := os.Open(s.filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var up core.URLPair
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		json.Unmarshal(scanner.Bytes(), &up)
		if up.Short == shortURL {
			return &core.URL{OriginalURL: up.Origin, Key: up.Short}, nil
		}
	}

	return nil, fmt.Errorf("не найдена пара url")

}

// InsertUser сохраняет нового пользователя или возвращает уже имеющегося в наличии
func (s fileStorage) InsertUser(Key int) (*core.User, error) {
	if Key <= 0 {
		user := &core.User{Key: s.getNextFreeKey(), WasUnathorized: true}
		s.users[user.Key] = user
		return user, nil
	}
	user, isExist := s.users[Key]
	if !isExist {
		user = &core.User{Key: Key, WasUnathorized: true}
		s.users[Key] = user
		return user, nil
	}
	return user, nil
}

// InsertUserURLPair cохраняет информацию о том, что пользователь сокращал URL, если такой информации ранее не было
func (s fileStorage) insertUserURLPair(shorten, origin string, user *core.User) error {
	URLPair := &core.UserURLPair{UserKey: user.Key, Short: shorten, Origin: origin}

	userURLPairs, isExist := s.userURLPairs[URLPair.UserKey]
	if !isExist {
		userURLPair := *URLPair
		new := make([]core.UserURLPair, 0, 10)
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

	fmt.Println("Хранится историй запросов пользователей на данный момент: ")
	for _, v := range s.userURLPairs {
		fmt.Println(*v)
	}

	return nil
}

func (s fileStorage) SelectUserByKey(Key int) (*core.User, error) {
	user, isExist := s.users[Key]
	if !isExist {
		return nil, fmt.Errorf("нет пользователя с ключом: %d", Key)
	}
	return user, nil
}

// SelectUserURL возвращает перечень соответствий между оригинальным и коротким адресом для конкретного пользователя
func (s fileStorage) SelectUserURLHistory(user *core.User) ([]core.UserURLPair, error) {
	if s.userURLPairs[user.Key] == nil {
		return nil, fmt.Errorf("нет истории")
	}
	return *s.userURLPairs[user.Key], nil
}

// getNextFreeKey возвращает ближайший свободный идентификатор пользователя
func (s fileStorage) getNextFreeKey() int {
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

func (s fileStorage) CloseConnection() {
	fmt.Println("Закрыто")
}

func (s fileStorage) DeleteURLs(URLs []core.URL) error {
	return nil
}

func (s fileStorage) CheckFile(filePath string) error {
	fileInfo, err := os.Stat(filePath)

	if errors.Is(err, os.ErrNotExist) {
		newFile, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
			return err
		}
		newFile.Close()
		fmt.Println("Создан файл: ", newFile.Name())
		return nil
	}
	fmt.Println("Файл уже есть: ", fileInfo.Name())
	return nil
}

func (s fileStorage) PingDB() error {
	return nil
}

type fileStorage struct {
	filePath     string
	users        map[int]*core.User
	userURLPairs map[int]*[]core.UserURLPair
}
