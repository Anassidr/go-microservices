package server

import (
	"context"
	"io"
	"time"

	"github.com/anassidr/go-microservices/currency/data"
	protos "github.com/anassidr/go-microservices/currency/protos/generated"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Currency is a gRPC server, it implements the methods defined by the CurrencyServer interface
type Currency struct {
	protos.UnimplementedCurrencyServer
	rates         *data.ExchangeRates
	log           hclog.Logger
	subscriptions map[protos.Currency_SubscribeRatesServer][]*protos.RateRequest
}

// NewCurrency creates a new Currency server
func NewCurrency(r *data.ExchangeRates, l hclog.Logger) *Currency {
	c := &Currency{rates: r, log: l, subscriptions: make(map[protos.Currency_SubscribeRatesServer][]*protos.RateRequest)}
	go c.handleUpdates()

	return c
}

func (c *Currency) handleUpdates() {
	ru := c.rates.MonitorRates(5 * time.Second)
	for range ru {
		c.log.Info("Got updated rates")

		// loop over subscribed clients
		for k, v := range c.subscriptions {
			// loop over subscribed rates
			for _, rr := range v {
				r, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
				if err != nil {
					c.log.Error("Unable to get updated rate", "base", rr.GetBase().String(), "destination", rr.GetDestination().String())
				}
				err = k.Send(
					&protos.StreamingRateResponse{
						Message: &protos.StreamingRateResponse_RateResponse{
							RateResponse: &protos.RateResponse{
								Base: rr.Base, Destination: rr.Destination, Rate: r},
						},
					},
				)

				if err != nil {
					c.log.Error("Unable to send updated rate", "base", rr.GetBase().String(), "destination", rr.GetDestination().String())

				}
			}

		}
	}
}

// GetRate implements the CurrencyServer GetRate method and returns the currency exchange rate for the two
// given currencies

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())

	if rr.Base == rr.Destination {
		// how to handle errors in grpc
		err := status.Newf(
			codes.InvalidArgument,
			"Base currency %s cannot be the same as the destination currency %s",
			rr.Base.String(),
			rr.Destination.String(),
		)
		err, wde := err.WithDetails(rr)
		if wde != nil {
			return nil, wde
		}
		return nil, err.Err()
	}

	rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
	if err != nil {
		return nil, err
	}

	return &protos.RateResponse{Base: rr.Base, Destination: rr.Destination, Rate: rate}, nil
}

// SubscribeRates implements the gRPC bidirection streaming method for the server

func (c *Currency) SubscribeRates(src protos.Currency_SubscribeRatesServer) error {

	// handle client message
	for {
		rr, err := src.Recv() // Recv is a blocking method that returns on client data

		// io.EOF signals that the client closed the connection
		if err == io.EOF {
			c.log.Info("Client has closed connection")
			break
		}

		// any other error means the transport btw the server and client is unavailable
		if err != nil {
			c.log.Error("Unable to read from client", "error", err)
			return err
		}
		c.log.Info("Handle client request", "request_base", rr.GetBase(), "request_dest", rr.GetDestination())

		rrs, ok := c.subscriptions[src]
		if !ok {
			rrs = []*protos.RateRequest{}
		}

		// check that the subscription does not exist
		var validationError *status.Status
		for _, v := range rrs {
			if v.Base == rr.Base && v.Destination == rr.Destination {
				validationError = status.Newf(
					codes.AlreadyExists,
					"Unable to subscribe for currency as subscription already exists",
				)
				// add the original request as metadata
				validationError, err = validationError.WithDetails(rr)
				if err != nil {
					c.log.Error("Unable to add metadata to error", "error", err)
					break
				}

				break
			}
		}

		// if a validation error return error and continue
		if validationError != nil {
			src.Send(
				&protos.StreamingRateResponse{
					Message: &protos.StreamingRateResponse_Error{
						Error: validationError.Proto(),
					},
				},
			)
			continue
		}
		// all ok
		rrs = append(rrs, rr)
		c.subscriptions[src] = rrs
	}
	return nil
}
