package handlers

import (
	"context"
	"net/http"

	"github.com/anassidr/go-microservices/data"
)

// defining this middleware function in order to avoid writing it inside very handler function
// code reuse
// The middleware function is responsible for performing validation of incoming product data
// and then passing the request down the middleware chain to the next handler function

// using a value receiver instead of pointer receiver since we do not need ot modify the Products object, only access its properties

func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := data.FromJSON(prod, r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)

			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		// validate the product
		errs := p.v.Validate(prod)
		if len(errs) != 0 {
			p.l.Println("[ERROR] validating product", errs)

			// return the validation messages as an array
			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}

//request object goes from client to server, through the middleware functions that modify the request (context or headers)
//response object  goes back, passing through the same chain of middlewares
//we are effectively building a pipeline
