package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "anass",
		Price: 10,
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
