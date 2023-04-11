package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/anassidr/go-microservices/product-api/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&prod)
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	//fetch products from datastore
	lp := data.GetProducts()
	//serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Products")
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

}

// defining this middleware function in order to avoid writing it inside very handler function
// code reuse
// The middleware function is responsible for performing validation of incoming product data
// and then passing the request down the middleware chain to the next handler function

type KeyProduct struct{}

// using a value receiver instead of pointer receiver since we do not need ot modify the Products object, only access its properties

func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(

		func(rw http.ResponseWriter, r *http.Request) {
			prod := data.Product{}

			err := prod.FromJSON(r.Body)
			if err != nil {
				p.l.Println("[ERROR] deserializing product", err)
				http.Error(rw, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
				return
			}
			//validate the product
			err = prod.Validate()
			if err != nil {
				p.l.Println("[ERROR] validating product", err)
				http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
				return
			}

			//add the product to the context
			ctx := context.WithValue(r.Context(), KeyProduct{}, prod) //Withvalue method on the context object is used to add a key-value pair to the context object
			req := r.WithContext(ctx)

			next.ServeHTTP(rw, req) //call the next handler in the middleware
		})
}

//request object goes from client to server, through the middleware functions that modify the request (context or headers)
//response object  goes back, passing through the same chain of middlewares
//we are effectively building a pipeline
