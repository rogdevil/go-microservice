package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func NewGoodBoi(l *log.Logger) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		l.Println("response from good boi")

		data, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprint(rw, string(data))
	})
}
