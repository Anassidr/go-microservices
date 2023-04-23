package data

import (
	"encoding/xml"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/go-hclog"
)

type ExchangeRates struct {
	log   hclog.Logger
	rates map[string]float64
}

func NewRates(l hclog.Logger) (*ExchangeRates, error) {
	er := &ExchangeRates{log: l, rates: map[string]float64{}}
	err := er.getRates()

	return er, err
}

func (e *ExchangeRates) GetRate(base string, dest string) (float64, error) {
	br, ok := e.rates[base]
	if !ok {
		return 0, fmt.Errorf("rate not found for currency %s", base)
	}

	dr, ok := e.rates[dest]
	if !ok {
		return 0, fmt.Errorf("rate not found for currency %s", dest)
	}
	return dr / br, nil
}

// MonitorRates checks the rates in the ECB API every interval and sends a message to the returned channel
// when there are changes
// Given that the CB API only returns data once a day, this functions simulates changes randomly

func (e *ExchangeRates) MonitorRates(interval time.Duration) chan struct{} {
	ret := make(chan struct{})

	go func() {
		ticker := time.NewTicker((interval))
		for {
			select {
			case <-ticker.C:
				for k, v := range e.rates {
					// random change in the rate; simulate the real-life fluctuations
					// change can be 10% of original value
					change := (rand.Float64() / 10)
					// is this a positive or negative change
					direction := rand.Intn(2)

					if direction == 0 {
						// new value will be min 90% of old
						change = 1 - change
					} else {
						// new value will be 110% of old
						change = 1 + change
					}
					// modify the rate
					e.rates[k] = v * change
				}
				// notify updates, this will block unless there is a listener on the other end
				// sending an empty struct to signal that an update has occured
				ret <- struct{}{}
			}
		}
	}()
	return ret
}

// Consuming the API of ECB

func (e *ExchangeRates) getRates() error {
	resp, err := http.DefaultClient.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")
	if err != nil {
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("expected error code 200, got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	md := &Cubes{}
	// NewDecoder wants an io Reader, we get it from resp.Body
	// Unmarsheling the XML data
	xml.NewDecoder(resp.Body).Decode(&md)

	for _, c := range md.CubeData {
		r, err := strconv.ParseFloat(c.Rate, 64)
		if err != nil {
			return err
		}
		e.rates[c.Currency] = r
	}
	e.rates["EUR"] = 1
	return nil
}

type Cubes struct {
	CubeData []Cube `xml:"Cube>Cube>Cube"`
}

type Cube struct {
	Currency string `xml:"currency,attr"`
	Rate     string `xml:"rate,attr"`
}
