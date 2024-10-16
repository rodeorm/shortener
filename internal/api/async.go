package api

import (
	"fmt"
	"sync"

	"github.com/rodeorm/shortener/internal/core"
	"github.com/rodeorm/shortener/internal/logger"
	"github.com/rodeorm/shortener/internal/repo"
)

func (q *Queue) Push(url []core.URL) error {
	var wg sync.WaitGroup

	for _, v := range url {
		wg.Add(1)
		go func() {
			q.ch <- &v
			wg.Done()
		}()
	}

	wg.Wait()
	return nil
}

func NewQueue(n int) *Queue {
	return &Queue{
		ch: make(chan *core.URL, n),
	}
}

type Queue struct {
	ch chan *core.URL
}

func (q *Queue) PopWait(n int) []core.URL {

	urls := make([]core.URL, 0)
	for i := 0; i < n; i++ {
		select {
		case val := <-q.ch:
			urls = append(urls, *val)
		default:
			continue
		}
	}
	return urls
}

type Worker struct {
	id        int
	batchSize int
	queue     *Queue
	storage   repo.Storager
}

func NewWorker(id int, queue *Queue, storage repo.Storager, batchSize int) *Worker {
	w := Worker{
		id:        id,
		queue:     queue,
		storage:   storage,
		batchSize: batchSize,
	}
	return &w
}

func (w *Worker) Loop() {
	logger.Log.Info(fmt.Sprintf("воркер #%d стартовал", w.id))

	for {
		urls := w.queue.PopWait(w.batchSize)

		if len(urls) == 0 {
			continue
		}
		err := w.storage.DeleteURLs(urls)
		if err != nil {
			logger.Log.Error(fmt.Sprintf("ошибка при работе воркера %v стартовал", err))
			continue
		}
		logger.Log.Info(fmt.Sprintf("воркер #%d удалил пачку урл %v", w.id, urls))
	}
}
