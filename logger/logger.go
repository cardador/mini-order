package logger

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		fmt.Printf("Method: %s | Path: %s | Duration: %s\n", r.Method, r.URL.Path, time.Since(start))
	}
}
