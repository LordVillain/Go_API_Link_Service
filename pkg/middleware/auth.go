package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authedHeader := r.Header.Get("Authorization")

		if authedHeader == "" {
			next.ServeHTTP(w, r)
			return
		}
		token := strings.TrimPrefix(authedHeader, "Bearer ")
		fmt.Println(token)
		next.ServeHTTP(w, r)
	})
}
