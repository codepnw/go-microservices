package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name: "Solo",
		Price: 1.25,
		SKU: "abc-def-hij",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
