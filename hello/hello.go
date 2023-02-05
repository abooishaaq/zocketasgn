package main

import (
	"net/http"
)

var response = []byte("Hello, World!")

func hello(w http.ResponseWriter, req *http.Request) {
	w.Write(response)
}

func main() {
	http.HandleFunc("/", hello)

	http.ListenAndServe(":8090", nil)
}
