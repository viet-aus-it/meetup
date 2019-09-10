package main

import (
	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"net/http"
)

type server struct {
	router *chi.Mux
	metrics *metrics
}

func main() {
	var s server
	s.router = chi.NewRouter()
	s.routes()
	s.metrics = newMetrics()
	err := s.metrics.register()
	if err != nil {
		logrus.WithError(err).Println("unable to register metrics")
	}

	logrus.
		WithError(http.ListenAndServe(":8484", s.router)).
		Panic()
}

// routes.go
func (s *server) routes() {
	s.router.Handle("/metrics", promhttp.Handler())
	s.router.HandleFunc("/dig", s.handleDig())
	s.router.HandleFunc("/close", s.handleClose())
}

// handler.go
func (s *server) handleDig() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(s.digNewHole()))
	}
}

func (s *server) handleClose() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(s.closeHole()))
	}
}

