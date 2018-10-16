package main

import (
	"strconv"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

//Temperature Code it the only documention you need.
type Temperature struct {
	Device string `json:"device"`
	Core int `json:"core"`
	Temp float32 `json:"temp"`
}

var save func(Temperature) error
var get func(int) ( []byte, error)

func main() {

	conn, err := NewInfluxClient()
	if err != nil {
		panic(err)
	}

	save = conn.save
	get = conn.get

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

	r.Body.Read(body)

	json.Unmarshal(body, &temp)

	if err := save(temp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
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
		w.Write([]byte(err.Error()))
		return
	}

	temperature, err := get(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(temperature)

}

