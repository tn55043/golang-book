package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func HomePageHandle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	fmt.Printf("Start at %v", start)

	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
	fmt.Printf("Completed in %v", time.Since(start))
}

func UsersHandle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	fmt.Printf("Start at %v", start)

	fmt.Fprintf(w, "Users Page")

	fmt.Printf("Completed in %v", time.Since(start))
}

func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", HomePageHandle).Methods("GET")
	r.HandleFunc("/users", UsersHandle).Methods("GET")
	r.Use(loggingMiddleware)
	return r
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("Start at %v", start)

		next.ServeHTTP(w, r)

		fmt.Printf("Completed in %v", time.Since(start))
	})
}

func main() {
	http.ListenAndServe(":3000", NewRouter())
}
