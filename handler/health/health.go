package health

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Health string

func (h Health) Handle() http.Handler {
	c := chi.NewRouter()
	c.Get(`/`, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type","text/plain")
		_, _ = fmt.Fprintf(w, string(h))
	})
	return c
}

func (h Health) Pattern() string {
	return "/health"
}





