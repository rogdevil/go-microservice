package handler

import (
	"log"
	"net/http"

	"github.com/rogdevil/data"
)

func NewProducts(l *log.Logger) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			GetProducts(rw, r, l)
			return
		}

		// catch all
		rw.WriteHeader(http.StatusMethodNotAllowed)
	})
}

func GetProducts(rw http.ResponseWriter, r *http.Request, l *log.Logger) {
	l.Println("fetching products ", r.RemoteAddr)
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshel JSON data", http.StatusInternalServerError)
		return
	}
}
