package main

import (
	"flag"
	"os"

	"github.com/rodeorm/shortener/internal/api"
	"github.com/rodeorm/shortener/internal/logger"
	"github.com/rodeorm/shortener/internal/repo"
)

// config выполняет первоначальную конфигурацию
func config() *api.Server {
	flag.Parse()

	// os.Setenv("SERVER_ADDRESS", "localhost:8080")
	// os.Setenv("BASE_URL", "http://tiny")
	// os.Setenv("FILE_STORAGE_PATH", "D:/file.txt")
	// os.Setenv("DATABASE_DSN", "postgres://app:qqqQQQ123@localhost:5432/shortener?sslmode=disable")

	var serverAddress, baseURL, fileStoragePath, databaseConnectionString string
	var workerCount, batchSize, queueSize int

	//Адрес запуска HTTP-сервера
	if *a == "" {
		serverAddress = os.Getenv("SERVER_ADDRESS")
		if serverAddress == "" {
			serverAddress = "localhost:8080"
		}
	} else {
		serverAddress = *a
	}

	//Базовый адрес результирующего сокращённого URL
	if *b == "" {
		baseURL = os.Getenv("BASE_URL")
		if baseURL == "" {
			baseURL = "http://localhost:8080"
		}
	} else {
		baseURL = *b
	}

	//Путь до файла
	if *f == "" {
		fileStoragePath = os.Getenv("FILE_STORAGE_PATH")
	} else {
		fileStoragePath = *f
	}

	//Строка подключения к БД
	if *d == "" {
		databaseConnectionString = os.Getenv("DATABASE_DSN")
	} else {
		databaseConnectionString = *d
	}

	if *w == "" {
		workerCount = 2
	}

	if *s == "" {
		batchSize = 3
	}

	if *q == "" {
		queueSize = 10
	}

	logger.Initialize("info")

	server := &api.Server{
		ServerAddress: serverAddress,
		Storage:       repo.NewStorage(fileStoragePath, databaseConnectionString),
		DeleteQueue:   api.NewQueue(queueSize),
		BaseURL:       baseURL,
		WorkerCount:   workerCount,
		BatchSize:     batchSize}

	return server
}
