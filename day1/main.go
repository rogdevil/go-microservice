package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("hello world")
	})

	http.HandleFunc("/goodboi", func(rw http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(rw, "morty its a error", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "hello mother fucker %s", data)
	})
	http.ListenAndServe("0.0.0.0:8080", nil)
}
