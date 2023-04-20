package data

import (
	"testing"

	"fmt"

	"github.com/hashicorp/go-hclog"
)

func TestNewRates(t *testing.T) {
	tr, err := NewRates(hclog.Default())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v", tr.rates)
}
