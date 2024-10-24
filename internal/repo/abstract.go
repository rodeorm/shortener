package repo

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/rodeorm/shortener/internal/core"
	"github.com/rodeorm/shortener/internal/logger"
	"go.uber.org/zap"
)

// Storager - абстрация для взаимодействия с хранилищем данных
type Storager interface {
	/*
		InsertURL принимает оригинальный URL, базовый урл для генерации коротких адресов и пользователя.
		Генерирует уникальный ключ для короткого адреса, сохраняет соответствие оригинального URL и ключа.

		Возвращает обновленный URL с соответствующим сокращенным URL, а также признаком того, что URL сократили ранее.
	*/
	InsertURL(URL, baseURL string, user *core.User) (*core.URL, error)

	// SelectOriginalURL возвращает URL на основании короткого
	SelectOriginalURL(shortURL string) (*core.URL, error)

	// InsertUser сохраняет нового пользователя или возвращает уже имеющегося в наличии, а также значение "отсутствие авторизации по переданному идентификатору"
	InsertUser(Key int) (*core.User, error)

	// SelectUserURLHistory возвращает перечень соответствий между оригинальным и коротким адресом для конкретного пользователя
	SelectUserURLHistory(user *core.User) ([]core.UserURLPair, error)

	// DeleteURLs массово помечает URL как удаленные. Успешно удалить URL может только пользователь, его создавший.
	DeleteURLs(URLs []core.URL) error

	// Ниже методы специфичные для СУБД
	// Можно, конечно, определить несколько интерфейсов (чтобы специфичный для СУБД включал в себя интерфейсы общие)
	// Но это как показала практика - это лишние трудозатраты, усложняющие читабельность кода в местах использования интерфейса, поэтому вернул как было

	// Close закрывает соединение
	Close()
	// Ping проверяет соединение
	Ping() error
}

// NewStorage определяет место для хранения данных
func NewStorage(filePath, dbConnectionString string) Storager {
	var storage Storager

	storage, err := InitPostgresStorage(dbConnectionString)
	if err == nil {
		return storage
	}

	if filePath != "" {
		storage, err = InitFileStorage(filePath)
		if err == nil {
			return storage
		}
	}
	storage = InitMemoryStorage()
	return storage
}

// InitMemoryStorage создает хранилище данных в оперативной памяти
func InitMemoryStorage() *memoryStorage {
	ots := make(map[string]string)
	sto := make(map[string]string)
	usr := make(map[int]*core.User)
	usrURL := make(map[int]*[]core.UserURLPair)
	storage := memoryStorage{originalToShort: ots, shortToOriginal: sto, users: usr, userURLPairs: usrURL}

	logger.Log.Info("Init storage",
		zap.String("Storage", "Memory storage"),
	)

	return &storage
}

// InitFileStorage создает хранилище данных на файловой системе
func InitFileStorage(filePath string) (*fileStorage, error) {
	usr := make(map[int]*core.User)
	usrURL := make(map[int]*[]core.UserURLPair)

	storage := fileStorage{filePath: filePath, users: usr, userURLPairs: usrURL}
	if err := сheckFile(filePath); err != nil {
		return nil, err
	}
	logger.Log.Info("Init storage",
		zap.String("Storage", "File storage"),
	)

	return &storage, nil

}

// InitPostgresStorage создает хранилище данных в БД на экземпляре Postgres
func InitPostgresStorage(connectionString string) (*postgresStorage, error) {
	db, err := sqlx.Open("pgx", connectionString)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}
	delQueue := make(chan string)
	storage := postgresStorage{DB: db, ConnectionString: connectionString, deleteQueue: delQueue, preparedStatements: map[string]*sqlx.Stmt{}}
	err = storage.createTables(ctx)
	if err != nil {
		return nil, err
	}

	err = storage.prepareStatements()
	if err != nil {
		return nil, err
	}

	logger.Log.Info("Init storage",
		zap.String("Storage", "PostgresStorage"),
	)

	return &storage, nil
}

func (s *postgresStorage) prepareStatements() error {

	nstmtSelectUser, err := s.DB.Preparex(`SELECT ID from Users WHERE ID = $1`)
	if err != nil {
		return err
	}

	nstmtInsertUser, err := s.DB.Preparex(`INSERT INTO Users (Name) VALUES ($1) RETURNING ID`)
	if err != nil {
		return err
	}

	nstmtSelectShortURL, err := s.DB.Preparex(`SELECT short from Urls WHERE original = $1`)
	if err != nil {
		return err
	}
	nstmtInsertURL, err := s.DB.Preparex(`INSERT INTO Urls (original, short, userID) SELECT $1, $2, $3`)
	if err != nil {
		return err
	}

	nstmtSelectOriginalURL, err := s.DB.Preparex(`SELECT original, isDeleted FROM Urls WHERE short = $1`)
	if err != nil {
		return err
	}

	nstmtSelectUserURLHistory, err := s.DB.Preparex(`SELECT original AS origin, short, userID AS userkey FROM Urls WHERE UserID = $1`)
	if err != nil {
		return err
	}

	nstmtDeleteURL, err := s.DB.Preparex(`UPDATE Urls SET isDeleted = true WHERE short = $1 AND userID = $2`)
	if err != nil {
		return err
	}

	// deleteURL UPDATE Urls SET isDeleted = true WHERE short = $1 AND userID = $2

	s.preparedStatements["SelectUser"] = nstmtSelectUser
	s.preparedStatements["InsertUser"] = nstmtInsertUser
	s.preparedStatements["SelectShortURL"] = nstmtSelectShortURL
	s.preparedStatements["InsertURL"] = nstmtInsertURL
	s.preparedStatements["SelectOriginalURL"] = nstmtSelectOriginalURL
	s.preparedStatements["SelectUserURLHistory"] = nstmtSelectUserURLHistory
	s.preparedStatements["DeleteURL"] = nstmtDeleteURL

	return nil
}
