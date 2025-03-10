package main

import (
	"flag"
)

var a, b, c, config, d, f, w, s, q, p, bs, t *string

func init() {

	//флаг -a, отвечающий за адрес запуска HTTP-сервера (переменная SERVER_ADDRESS)
	a = flag.String("a", "", "SERVER_ADDRESS")
	//флаг -b, отвечающий за базовый адрес результирующего сокращённого URL (переменная BASE_URL)
	b = flag.String("b", "", "BASE_URL")
	//флаг -bs, отвечающий за размер пачки для удаления
	bs = flag.String("bs", "", "BUTCH_SIZE")
	//флаг -c, отвечающий за имя файла конфигурации
	c = flag.String("c", "", "CONFIG")
	//флаг -config, отвечающий за имя файла конфигурации
	config = flag.String("config", "", "CONFIG")
	//флаг -d, отвечающий за строку подключения к БД (переменная DATABASE_DSN)
	d = flag.String("d", "", "DATABASE_DSN")
	//флаг -f, отвечающий за путь до файла с сокращёнными URL (переменная FILE_STORAGE_PATH)
	f = flag.String("f", "", "FILE_STORAGE_PATH")
	//флаг -p, отвечающий за тип профилирования
	p = flag.String("p", "", "PROFILE_TYPE")
	//флаг -q, отвечающий за размер очереди для удаления
	q = flag.String("q", "", "DELETE_QUEUE_SIZE")
	//флаг -s, отвечающий за включение https
	s = flag.String("s", "", "HTTPS")
	//флаг -t, отвечающий за строковое представление бесклассовой адресации (CIDR)
	t = flag.String("t", "", "TRUSTED_SUBNET")
	//флаг -w, отвечающий за число воркеров для удаления
	w = flag.String("w", "", "WORKER_COUNTS")

	logInit()

	printBuildInfo()
}
