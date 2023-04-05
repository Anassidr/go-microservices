package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/anassidr/go-microservices/handlers"
	"golang.org/x/net/context"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Goroutine : execute function concurrently with other functions
	// Goroutines are different from traditional OS threads because they are managed by the GO runtime. Makes them more light weight.
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal) //create a new channel of signals, receive notifications for interrupt signals
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan //wait for a signal to be sent to sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc) //wait until the requests that are currently handled by the server are completed and then shut down
}
