package main

import (
	"fmt"
	"net/http"
	"time"
)

func logMiddlweare(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Logging from middlware")
		f(w, r)
	}
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /persons", logMiddlweare(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 30)
		w.Write([]byte("Test\n"))
	}))
	router.HandleFunc("GET /persons/items", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test concurrency\n"))
	})
	http.ListenAndServe("localhost:8080", router)
}
