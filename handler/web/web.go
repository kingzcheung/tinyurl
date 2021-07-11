package web

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golang/groupcache"
	"github.com/kingzcheung/tinyurl/config"
	"github.com/kingzcheung/tinyurl/core"
	"github.com/kingzcheung/tinyurl/static"
	"github.com/sirupsen/logrus"
	"net/http"
	"path"
)

var contentTypes = map[string]string{
	".html": "text/html; charset=UTF-8",
	".js":   "text/javascript; charset=UTF-8",
	".css":  "text/css; charset=utf-8",
	".png":  "image/png",
	".jpg":  "image/jpeg",
	".jpeg": "image/jpeg",
	".svg":  "image/svg+xml",
}

type Server struct {
	Cache  *groupcache.Group
	Config *config.Config
}

func NewServer(cache *groupcache.Group, config *config.Config) *Server {
	return &Server{Cache: cache, Config: config}
}

func (s *Server) Handle() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)
	r.NotFound(handlerIndex(s.Config.AppMode))
	r.Get(`/robots.txt`, robots())
	r.Get("/r/{id}", Redirect(s.Cache))
	//r.Get(`/callback`, oauth.AccessTokenCB(s.Config))

	return r
}

func handlerIndex(mode string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// dev debug
		//if mode == "dev" {
		//	frontEndUrl := "http://127.0.0.1:3000"
		//	remote, _ := url.Parse(frontEndUrl)
		//	proxy := httputil.NewSingleHostReverseProxy(remote)
		//	proxy.ServeHTTP(w, r)
		//	return
		//}

		uri := r.URL.Path
		if uri == "/" {
			uri = "index.html"
		}

		filePath := path.Join("dist", uri)

		file, err := static.WebDict.ReadFile(filePath)
		if err != nil {
			logrus.WithField("filePath", filePath).Error(err.Error())
			http.Error(w, "404 not found", http.StatusNotFound)
			return
		}
		ext := path.Ext(filePath)
		var contentType, ok = contentTypes[ext]
		if !ok {
			contentType = "text/plain"
		}

		w.Header().Set("Content-Type", contentType)
		_, _ = w.Write(file)
	}
}

func robots() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("User-agent: * \n"))
		_, _ = w.Write([]byte("Disallow: /"))
	}
}

func (s *Server) Pattern() string {
	return "/"
}

func Redirect(cache *groupcache.Group) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		hashID := chi.URLParam(r, "id")
		if hashID == "" {
			http.Error(w, "404 NotFound", http.StatusNotFound)
			return
		}

		var data []byte
		err := cache.Get(nil, hashID, groupcache.AllocatingByteSliceSink(&data))
		if err != nil {
			http.Error(w, "404 NotFound", http.StatusNotFound)
			return
		}
		var u core.Url
		_ = json.Unmarshal(data, &u)

		if u.URL == "" {
			http.Error(w, "404 NotFound", http.StatusNotFound)
			return
		}

		//tracking.Pool.Process(tracking.NewService(r, db, u.UrlID))

		//_ = tracking.NewService(r, h.Database).Track(u.UrlID)

		http.Redirect(w, r, u.URL, http.StatusFound)
	}
}
