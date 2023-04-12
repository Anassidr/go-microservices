package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/anassidr/go-microservices/product-api/data"
)

// defining this middleware function in order to avoid writing it inside very handler function
// code reuse
// The middleware function is responsible for performing validation of incoming product data
// and then passing the request down the middleware chain to the next handler function

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
