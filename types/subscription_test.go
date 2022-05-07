package types

import (
	"encoding/json"
	"testing"
)

// TestExampleSubscriptions tests whether the API documentation's example subscriptions unmarshal correctly to the Subscriptions type
func TestExampleSubscriptions(t *testing.T) {
	subscriptions := &Subscriptions{}
	if err := json.Unmarshal(exampleSubscriptions, subscriptions); err != nil {
		t.Fatal(err)
	}
}
