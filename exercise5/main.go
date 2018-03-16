package main

import (
	"encoding/json"
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
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
		TempMin  int     `json:"temp_min"`
		TempMax  int     `json:"temp_max"`
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
	fmt.Printf("Start at %v\n", start)

	vars := mux.Vars(r)
	name := vars["name"]

	res, _ := http.Get("http://localhost:8882/api/v1/weather/" + name)

	wstub := new(WeatherStub)
	json.NewDecoder(res.Body).Decode(wstub)

	fmt.Fprintf(w, "%s\n", wstub.Name)
	fmt.Fprintf(w, "%.2fc %s", wstub.Main.Temp, wstub.Weather[0].Description)
	fmt.Printf("Completed in %v\n", time.Since(start))
}

func WeatherAllHandle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	fmt.Printf("Start at %v\n", start)

	country := [5]string{"hobart", "newyork", "kupang", "nairobi", "bangkok"}
	for i := 0; i < len(country); i++ {
		res, err := http.Get("http://localhost:8882/api/v1/weather/" + country[i])
		if err != nil {
			fmt.Println(err)

			return
		}

		wstub := new(WeatherStub)
		json.NewDecoder(res.Body).Decode(wstub)

		fmt.Fprintf(w, "%s\n", wstub.Name)
		fmt.Fprintf(w, "%2.fc %s\n\n", wstub.Main.Temp, wstub.Weather[0].Description)
	}

	fmt.Printf("Completed in %v\n", time.Since(start))
}

func NewRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/weather/all", WeatherAllHandle).Methods("GET")
	r.HandleFunc("/weather/{name}", WeatherHandle).Methods("GET")
	r.Use(loggingMiddleware)
	return r
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("Start at %v\n", start)

		next.ServeHTTP(w, r)

		fmt.Printf("Completed in %v\n", time.Since(start))
	})
}

func main() {
	http.ListenAndServe(":3000", NewRouter())
}
