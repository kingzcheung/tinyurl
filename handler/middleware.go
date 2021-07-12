package handler

import (
	"github.com/alexedwards/scs/v2"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func CheckLogin(session *scs.SessionManager) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			username := session.GetString(r.Context(), "username")
			if username == "" {
				logrus.Error("The action requires login.")
				contentType := r.Header.Get("Content-Type")
				if strings.HasPrefix(contentType, "application/json") {
					w.WriteHeader(400)
					w.Header().Set("Content-Type", "application/json")
					_, _ = w.Write([]byte(`{"error":"Please log in"}`))
				} else {
					http.Error(w, "404 not found", 404)
				}
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
