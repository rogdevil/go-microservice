package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle Hello requests")
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		h.l.Println("Error reading request body", err)

		http.Error(rw, "morty its a error", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "hello %s", data)
}
