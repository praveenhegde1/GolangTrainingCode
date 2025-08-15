package main

import (
	"net/http"

	"go.uber.org/zap"
)

var logger, _ = zap.NewProduction()

func handler(w http.ResponseWriter, r *http.Request) {
	logger.Info("Received request",
		zap.String("method", r.Method),
		zap.String("url", r.URL.String()),
		zap.String("remoteAddr", r.RemoteAddr),
	)
	w.Write([]byte("Hello, World!"))
}

func main() {
	defer logger.Sync()
	http.HandleFunc("/", handler)
	logger.Info("Starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}
