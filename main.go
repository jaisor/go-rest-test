package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	logger := log.New(os.Stdout, "Jaisor-Go-Rest", log.LstdFlags)
	env := &Env{log: logger}

	logger.Println("Jaisor : Go REST Test")

	m := mux.NewRouter().StrictSlash(true)
	m.Use(commonMiddleware)

	m.HandleFunc("/ping", env.healthCheck).Methods("GET")
	m.HandleFunc("/dump", env.dumpRequest)
	m.HandleFunc("/", env.index)

	port := getEnv("IP_ADDR", ":10000")
	logger.Printf("Attempting to listen on port: %s", port)
	log.Fatal(http.ListenAndServe(port, m))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// Env - DI of environment
type Env struct {
	log *log.Logger
}

func (env *Env) index(w http.ResponseWriter, r *http.Request) {

	health := map[string]interface{}{
		"go": "gone",
	}

	json.NewEncoder(w).Encode(health)
}

func (env *Env) healthCheck(w http.ResponseWriter, r *http.Request) {

	health := map[string]interface{}{
		"go":  "true",
		"now": time.Now(),
	}

	json.NewEncoder(w).Encode(health)
}

func (env *Env) dumpRequest(w http.ResponseWriter, r *http.Request) {

	env.log.Printf("Responding to '%s' request from %s", r.RequestURI, r.RemoteAddr)

	response := map[string]interface{}{
		"scucess": true,
		"method":  r.Method,
		"headers": r.Header,
		"body":    r.Body,
		"url":     r.URL,
		"now":     time.Now(),
	}

	json.NewEncoder(w).Encode(response)
}
