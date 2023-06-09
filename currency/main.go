package main

import (
	"net"
	"os"

	"github.com/anassidr/go-microservices/currency/data"
	protos "github.com/anassidr/go-microservices/currency/protos/generated"
	"github.com/anassidr/go-microservices/currency/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	rates, err := data.NewRates(log)

	if err != nil {
		log.Error("Unable to generate rates", "error", err)
		os.Exit(1)
	}

	gs := grpc.NewServer()
	cs := server.NewCurrency(rates, log)

	protos.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs) // disable in production

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}
	gs.Serve(l)
}
