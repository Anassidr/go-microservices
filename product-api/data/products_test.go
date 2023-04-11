package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "anass",
		Price: 10,
		SKU:   "abs-zje-qlf",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
