package handler

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/rogdevil/data"
)

func NewProducts(l *log.Logger) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			GetProducts(rw, r, l)
			return
		}

		if r.Method == http.MethodPost {
			AddProduct(rw, r, l)
			return
		}

		if r.Method == http.MethodPut {
			regExp := regexp.MustCompile(`/([0-9]+)`)
			matches := regExp.FindAllStringSubmatch(r.URL.Path, -1)

			if len(matches) != 1 {
				l.Println("Got more them one match of id")
				http.Error(rw, "Invalid URL", http.StatusBadRequest)
				return
			}

			l.Println(matches[0])

			if len(matches[0]) != 2 {
				l.Println("Got more them one sub match of id", matches[0])
				http.Error(rw, "Invalid URL", http.StatusBadRequest)
				return
			}

			prodId, err := strconv.Atoi(matches[0][1])

			if err != nil {
				l.Println("Unable to parse the id properly")
				http.Error(rw, "Invalid URL", http.StatusInternalServerError)
				return
			}

			UpdateProduct(rw, r, l, prodId)
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

func AddProduct(rw http.ResponseWriter, r *http.Request, l *log.Logger) {
	l.Println("adding product", r.RemoteAddr)
	lp := &data.Product{}
	err := lp.FromJSON(r.Body)

	if err != nil {
		l.Println("there was a error add data:", err)
		http.Error(rw, "Unable to store data", http.StatusBadRequest)
		return
	}
	data.AddProduct(lp)
}

func UpdateProduct(rw http.ResponseWriter, r *http.Request, l *log.Logger, prodId int) {
	l.Println("updating product", r.RemoteAddr)
	lp := &data.Product{}
	err := lp.FromJSON(r.Body)

	if err != nil {
		l.Println("there was a error update data :", err)
		http.Error(rw, "Unable to store data", http.StatusBadRequest)
		return
	}
	err = data.UpdateProduct(lp, prodId)

	if err == data.ErrProductNotFound {
		l.Println("got error updating product:", err)
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Unable to updat the product data", http.StatusInternalServerError)
		return
	}
}
