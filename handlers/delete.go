package handlers

import (
	"net/http"
	"nic-microservices/data"
	"strconv"

	"github.com/gorilla/mux"
)

//swagger:route DELETE /{id} products deleteProduct
// Deletes a product by the id
// responses:
//	201: noContent
// 	404: errorResponse
// 	501: errorResponse
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	p.l.Println("Handle delete Product", id)

	err := data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}