package header

import (
	"context"
	"net/http"
)

// https://www.todobackend.com/
// https://github.com/savaki/todo-backend-gin/blob/master/todo.go

func withAPIKey(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")

		if !isValidAPIKey(key) {
			respondErr(w, r, http.StatusUnauthorized, "invalidAPI key")
			return
		}
		ctx := context.WithValue(r.Context(),
			contextKeyAPIKey, key)
		fn(w, r.WithContext(ctx))
	}
}

func t(w http.ResponseWriter, req *http.Request) {
	req.Header.Set("Authorization", "Bearer abc123")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")

	r.Header.Get("User-Agent")
}
