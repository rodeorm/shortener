package api

import (
	"log"
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/gorilla/mux"

	"github.com/rodeorm/shortener/internal/api/middleware"
	"github.com/rodeorm/shortener/internal/repo"
)

// ServerStart запускает веб-сервер
func ServerStart(s *Server) error {

	defer s.Storage.CloseConnection()
	defer close(s.DeleteQueue.ch)

	r := mux.NewRouter()
	r.HandleFunc("/", s.RootHandler).Methods(http.MethodPost)
	r.HandleFunc("/ping", s.PingDBHandler).Methods(http.MethodGet)
	r.HandleFunc("/{URL}", s.RootURLHandler).Methods(http.MethodGet)

	r.HandleFunc("/api/shorten", s.APIShortenHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/user/urls", s.APIUserGetURLsHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/user/urls", s.APIUserDeleteURLsHandler).Methods(http.MethodDelete)
	r.HandleFunc("/api/shorten/batch", s.APIShortenBatchHandler).Methods(http.MethodPost)

	r.HandleFunc("/", s.BadRequestHandler)
	r.Use(middleware.WithZip, middleware.WithLog)

	pprofRouter := r.PathPrefix("/debug/pprof/").Subrouter()
	pprofRouter.HandleFunc("/", pprof.Index)
	pprofRouter.HandleFunc("/cmdline", pprof.Cmdline)
	pprofRouter.HandleFunc("/profile", pprof.Profile)
	pprofRouter.HandleFunc("/symbol", pprof.Symbol)
	pprofRouter.HandleFunc("/trace", pprof.Trace)

	srv := &http.Server{
		Handler:      r,
		Addr:         s.ServerAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	for i := 0; i < s.WorkerCount; i++ {
		w := NewWorker(i, s.DeleteQueue, s.Storage, s.BatchSize)
		go w.Loop()
	}

	log.Fatal(srv.ListenAndServe())

	return nil
}

type Server struct {
	ServerAddress            string
	BaseURL                  string
	DatabaseConnectionString string

	WorkerCount int
	BatchSize   int
	ProfileType int

	Storage     repo.Storager
	DeleteQueue *Queue
}
