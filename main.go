package main

import (
	"strconv"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Temperature struct {
	ID int
	Temperature float32
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/temperature", TemperaturePutHandler).Methods("PUT", "POST")
	r.HandleFunc("/temperature/{id}", TemperatureGetHandler).Methods("GET")

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

// TemperaturePutHandler ...
func TemperaturePutHandler(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)

	var temp Temperature

	w.WriteHeader(http.StatusOK)
	r.Body.Read(body)

	json.Unmarshal(body, &temp)

	if err := save(temp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println(string(body))
}

// TemperatureGetHandler ...
func TemperatureGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	temperature, err := get(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(temperature))

}

