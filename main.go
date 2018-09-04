package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

// HomeHandler ...
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	body := []byte{}

	w.WriteHeader(http.StatusOK)
	r.Body.Read(body)

	fmt.Println(body)
}

