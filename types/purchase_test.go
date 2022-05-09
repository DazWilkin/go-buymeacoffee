package types

import (
	"encoding/json"
	"testing"
)

// TestExamplePuchases tests whether the API documentation's example subscriptions unmarshal correctly to the Purchases type
func TestExamplePuchases(t *testing.T) {
	purchases := &Purchases{}
	if err := json.Unmarshal(ExamplePurchases, purchases); err != nil {
		t.Fatal(err)
	}
}

// TestExamplePurchase tests whether the API documentation's example purchase unmarshals correctly
func TestExamplePurchase(t *testing.T) {
	purchase := &Purchase{}
	if err := json.Unmarshal(ExamplePurchase, purchase); err != nil {
		t.Fatal(err)
	}
}
