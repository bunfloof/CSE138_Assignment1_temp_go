package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path == "/hello" && r.Method == http.MethodGet:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "world"})

	case strings.HasPrefix(r.URL.Path, "/hello/") && r.Method == http.MethodPost:
		name := strings.TrimPrefix(r.URL.Path, "/hello/")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Hi, " + name + "."})

	case r.URL.Path == "/test" && r.Method == http.MethodGet:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "test is successful"})

	case r.URL.Path == "/test" && r.Method == http.MethodPost:
		msg := r.URL.Query().Get("msg")
		if msg == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Request"))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": msg})

	default:
		if (r.URL.Path == "/hello" && r.Method != http.MethodGet) ||
			(strings.HasPrefix(r.URL.Path, "/hello/") && r.Method != http.MethodPost) {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method Not Allowed"))
		} else {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Not Found"})
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8090", nil)
}
