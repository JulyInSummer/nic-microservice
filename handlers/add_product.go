package handlers

import (
	"net/http"

	"nj_microservices/data"
)

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle -> POST")

	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(prod)
	p.l.Printf("Product: %#v", prod)
}
