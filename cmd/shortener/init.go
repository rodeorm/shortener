package main

import (
	"flag"
)

var a, b, f, d, w, s, q, p *string

func init() {
	//флаг -a, отвечающий за адрес запуска HTTP-сервера (переменная SERVER_ADDRESS)
	a = flag.String("a", "", "SERVER_ADDRESS")
	//флаг -b, отвечающий за базовый адрес результирующего сокращённого URL (переменная BASE_URL)
	b = flag.String("b", "", "BASE_URL")
	//флаг -f, отвечающий за путь до файла с сокращёнными URL (переменная FILE_STORAGE_PATH)
	f = flag.String("f", "", "FILE_STORAGE_PATH")
	//флаг -d, отвечающий за строку подключения к БД (переменная DATABASE_DSN)
	d = flag.String("d", "", "DATABASE_DSN")
	//флаг -w, отвечающий за число воркеров для удаления
	w = flag.String("w", "", "WORKER_COUNTS")
	//флаг -s, отвечающий за размер пачки для удаления
	s = flag.String("s", "", "BUTCH_SIZE")
	//флаг -q, отвечающий за размер очереди для удаления
	q = flag.String("q", "", "DELETE_QUEUE_SIZE")
	//флаг -p, отвечающий за тип профилирования
	p = flag.String("p", "", "PROFILE_TYPE")
}

// Типы профилирования приложения
const (
	noneProfile   = iota // Нет профилирования
	baseProfile          // Профилирование в файл base
	resultProfile        // Профилирование в файл result
)
