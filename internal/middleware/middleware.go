package middleware

import (
	"log"
	"mail-phone-auth/internal/api/response"
	"mail-phone-auth/pkg/jwt"
	"net/http"
	"strings"
)

type Middleware struct {
	jwt *jwt.JWT
}

func New(jwt *jwt.JWT) *Middleware {
	return  &Middleware{
		jwt: jwt,
	}
}

func (m *Middleware) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		origin := r.Header.Get("Origin")
		if origin == "" {
			next.ServeHTTP(w, r)
			return
		}

		header := w.Header()
		header.Set("Access-Control-Allow-Origin", origin)
		header.Set("Access-Control-Allow-Credentials", "true")
		next.ServeHTTP(w, r)

	})
}
func (m *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authPaths := []string{
			"/api/user",
			"/api/file",
			"/api/role",
		}

		for _, path := range authPaths {
			if strings.HasPrefix(r.URL.Path, path) {
				authHeader := r.Header.Get("Authorization")
				if authHeader == "" {
					log.Println("Unauthorized")					
					response.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}
				const prefix = "Bearer "
				if !strings.HasPrefix(authHeader, prefix) {		
					log.Println("Unauthorized")					
					response.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}
				tokenString := authHeader[len(prefix):]
				_, err := m.jwt.ParseToken(tokenString)
				if err != nil {
					log.Println("Unauthorized")		
					response.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}				
			}
		}

		next.ServeHTTP(w, r)
	})
}
