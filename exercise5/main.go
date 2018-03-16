package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type WeatherStub struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     int `json:"temp"`
		Pressure int `json:"pressure"`
		Humidity int `json:"humidity"`
		TempMin  int `json:"temp_min"`
		TempMax  int `json:"temp_max"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

func WeatherHandle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	fmt.Printf("Start at %v", start)

	vars := mux.Vars(r)
	name := vars["name"]
	
	/*
	if vars["name"] == "" {
		name = "World"
	}
	*/

	fmt.Fprintf(w, "%s\n", name)
	fmt.Fprintf(w, "14c shower rain")
	fmt.Printf("Completed in %v", time.Since(start))
}

func WeatherAllHandle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	fmt.Printf("Start at %v", start)

	//

	fmt.Printf("Completed in %v", time.Since(start))
}

func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/weather/{name}", WeatherHandle).Methods("GET")
	r.HandleFunc("/weather/all", WeatherAllHandle).Methods("GET")
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
