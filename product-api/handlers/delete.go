package handlers

import (
	"net/http"

	"github.com/codepnw/microservices/api/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Return a list of products
// responses:
//	201: noContent

func (p *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := getProductID(r)

	p.l.Println("handle delete product", id)

	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)

		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
