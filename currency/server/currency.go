package server

import (
	"context"
	"io"
	"time"

	"github.com/anassidr/go-microservices/currency/data"
	protos "github.com/anassidr/go-microservices/currency/protos/generated"
	"github.com/hashicorp/go-hclog"
)

// Currency is a gRPC server, it implements the methods defined by the CurrencyServer interface
type Currency struct {
	protos.UnimplementedCurrencyServer
	rates *data.ExchangeRates
	log   hclog.Logger
}

// NewCurrency creates a new Currency server
func NewCurrency(r *data.ExchangeRates, l hclog.Logger) *Currency {
	return &Currency{rates: r, log: l}

}

// GetRate implements the CurrencyServer GetRate method and returns the currency exchange rate for the two
// given currencies

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())
	rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
	if err != nil {
		return nil, err
	}

	return &protos.RateResponse{Rate: rate}, nil
}

func (c *Currency) SubscribeRates(src protos.Currency_SubscribeRatesServer) error {

	go func() {
		for {
			rr, err := src.Recv()

			if err == io.EOF {
				c.log.Info("Client has closed connection")
				break
			}
			if err != nil {
				c.log.Error("Unable to read from client", "error", err)
				break
			}
			c.log.Info("Handle client request", "request", rr)
		}
	}()

	for {
		err := src.Send(&protos.RateResponse{Rate: 12.1})
		if err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
	}
}
