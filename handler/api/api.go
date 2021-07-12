package api

import (
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kingzcheung/tinyurl/core"
	"github.com/kingzcheung/tinyurl/handler"
	"github.com/kingzcheung/tinyurl/handler/api/url"
	"github.com/speps/go-hashids"
	"net/http"
)

type Server struct {
	urlStore  core.UrlStore
	userStore core.UserStore
	hashid    *hashids.HashID
	session   *scs.SessionManager
}

func NewServer(
	urlStore core.UrlStore,
	userStore core.UserStore,
	hashid *hashids.HashID,
	session *scs.SessionManager,
) *Server {
	return &Server{urlStore: urlStore, userStore: userStore, hashid: hashid, session: session}
}

func (s *Server) Handle() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)

	r.With(handler.CheckLogin(s.session)).Post(`/shorten`, url.Shorten(s.urlStore, s.hashid))
	r.Get(`/urls`, url.List(s.urlStore))
	//r.Get(`/urls/{url_id}/analytics`, url.Analytics(s.Config, s.Database))
	return r
}

func (s *Server) Pattern() string {
	return "/api"
}
