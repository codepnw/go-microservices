package handlers

import (
	"net/http"

	"github.com/codepnw/microservices/api/data"
)

// swagger:route GET /products products listProducts
// Return a list of products
// responses:
//	200: productsResponse

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	lp := data.GetProducts()
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
