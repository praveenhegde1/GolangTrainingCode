package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func handler(w http.ResponseWriter, r *http.Request) {
	log.WithFields(logrus.Fields{
		"method":     r.Method,
		"url":        r.URL.String(),
		"remoteAddr": r.RemoteAddr,
	}).Info("Received request")
	w.Write([]byte("Hello, World!"))
}

func main() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.DebugLevel)

	http.HandleFunc("/", handler)
	log.Info("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
