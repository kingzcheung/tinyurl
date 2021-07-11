package auth

import (
	"encoding/json"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/kingzcheung/tinyurl"
	"github.com/kingzcheung/tinyurl/config"
	"github.com/kingzcheung/tinyurl/core"
	"github.com/kingzcheung/tinyurl/static"
	"github.com/matthewhartstonge/argon2"
	"github.com/sirupsen/logrus"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
)

type Server struct {
	userStore core.UserStore
	session   *scs.SessionManager
	config    *config.Config
}

func NewServer(userStore core.UserStore, session *scs.SessionManager, config *config.Config) *Server {
	return &Server{userStore: userStore, session: session, config: config}
}

func (s *Server) Handle() http.Handler {
	r := chi.NewRouter()

	r.Post(`/`, handleLogin(s.session, s.userStore))
	r.Get(`/`, showLogin(s.config, s.session))
	return r
}

func showLogin(cnf *config.Config, manager *scs.SessionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// dev debug
		//frontEndUrl := "http://127.0.0.1:3000"
		//remote, _ := url.Parse(frontEndUrl)
		//proxy := httputil.NewSingleHostReverseProxy(remote)
		//proxy.ServeHTTP(w, r)
		//return

		loginText, err := static.WebDict.ReadFile(path.Join("dist", "login.html"))
		if err != nil {
			logrus.Error("login.html not found.")
			http.Error(w, "server error", 500)
		}
		tmp, _ := template.New("login").Parse(string(loginText))

		csrf := tinyurl.NewCsrf(cnf.Server.CsrfToken, manager)
		csrfToken := csrf.CreateToken(r.Context())

		_ = tmp.Execute(w, map[string]interface{}{
			"csrf_token": csrfToken,
		})

	}
}

func (s *Server) Pattern() string {
	return "/login"
}

//Login login
func handleLogin(manager *scs.SessionManager, userStore core.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "log in error", 400)
			return
		}
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		if strings.TrimSpace(username) == "" || strings.TrimSpace(password) == "" {
			http.Error(w, "username or password required", 400)
			return
		}

		user := userStore.FindUser(r.Context(), username)
		if user.UserID == 0 {
			http.Error(w, "user not found.", 400)
			return
		}
		ok, err := argon2.VerifyEncoded([]byte(password), []byte(user.Password))
		if err != nil {
			logrus.Errorf("verify error: %v", err)
			http.Error(w, "Please enter the correct username or password", 400)
			return
		}
		if !ok {
			http.Error(w, "Please enter the correct username or password", 400)
			return
		}

		manager.Put(r.Context(), "username", user.Username)

		http.Redirect(w, r, "/", http.StatusFound)
	}
}

//parseJson 从body中解析JSON
func parseJson(body io.ReadCloser, i interface{}) error {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, i)
}

func ErrJson(w http.ResponseWriter, v interface{}, status int) {
	w.WriteHeader(status)
	data, _ := json.Marshal(v)
	_, _ = w.Write(data)
}

func Json(w http.ResponseWriter, v interface{}) {
	ErrJson(w, v, http.StatusOK)
}
