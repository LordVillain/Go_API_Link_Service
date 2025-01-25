package middleware

import (
	"context"
	"go/adv-demo/configs"
	"go/adv-demo/pkg/jwt"
	"net/http"
	"strings"
)

type key string

const (
	ContextEmailKey key = "ContextEmailKey"
)

func WriteUnauthed(w http.ResponseWriter)  {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authedHeader, "Bearer "){
			WriteUnauthed(w)
			return
		}
		token := strings.TrimPrefix(authedHeader, "Bearer ")
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		if !isValid {
			WriteUnauthed(w)
			return
		}
		
		r.Context()
		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
