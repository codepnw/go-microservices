package handlers

import (
	"net/http"

	"github.com/codepnw/microservices/api/data"
)

// swagger:route POST /products products createProduct
// Create a new product
// responses:
//	200: productResponse

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}

	prod = r.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(prod)
}
