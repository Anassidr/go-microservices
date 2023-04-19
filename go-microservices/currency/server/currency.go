package server

import (
	"context"

	protos "github.com/anassidr/go-microservices/currency/protos/currency"
	"github.com/hashicorp/go-hclog"
)

// Currency is a gRPC server, it implements the methods defined by the CurrencyServer interface
type Currency struct {
	log hclog.Logger
}

// NewCurrency creates a new Currency server
func NewCurrency(l hclog.Logger) Currency {
	return Currency{l}

}

// GetRate implements the CurrencyServer GetRate method and returns the currency exchange rate for the two
// given currencies

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())
	return &protos.RateResponse{Rate: 0.5}, nil
}

func (c *Currency) mustEmbedUnimplementedCurrencyServer() {}
