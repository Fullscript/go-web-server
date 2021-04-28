package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s: ok", serverID)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s: responding from the go-web-server 👋", serverID)
}

func Keyword(w http.ResponseWriter, r *http.Request) {
	phrase, _ := base64.StdEncoding.DecodeString("d2VsbG5lc3M=")

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s: %s", serverID, phrase)
}
