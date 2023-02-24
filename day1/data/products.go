package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products is a collection of Product
type Products []*Product

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// GetProducts returns a list of products
func GetProducts() Products {
	return productList
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Product 1",
		Description: "Product 1 description",
		Price:       1.99,
		SKU:         "PRD-1",
		CreatedOn:   time.Now().Format("2006-01-02 15:04:05"),
		UpdatedOn:   time.Now().Format("2006-01-02 15:04:05"),
		DeletedOn:   "",
	},
	{
		ID:          2,
		Name:        "Product 2",
		Description: "Product 2 description",
		Price:       2.99,
		SKU:         "PRD-2",
		CreatedOn:   time.Now().Format("2006-01-02 15:04:05"),
		UpdatedOn:   time.Now().Format("2006-01-02 15:04:05"),
		DeletedOn:   "",
	},
}
