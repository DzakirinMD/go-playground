package main

import (
	"encoding/json"
	"go_playground/mini-project/weather-tracker/api"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/weather/", getWeather)

	//http.HandleFunc("/weather/",
	//	func(w http.ResponseWriter, r *http.Request) {
	//		city := strings.SplitN(r.URL.Path, "/", 3)[2]
	//		data, err := api.Query(city)
	//		if err != nil {
	//			http.Error(w, err.Error(), http.StatusInternalServerError)
	//			return
	//		}
	//		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//		json.NewEncoder(w).Encode(data)
	//	})

	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from go!\n"))
}

func getWeather(w http.ResponseWriter, r *http.Request) {
	city := strings.SplitN(r.URL.Path, "/", 3)[2]
	data, err := api.Query(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}
