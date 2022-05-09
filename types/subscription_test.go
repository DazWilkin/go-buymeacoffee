package types

import (
	"encoding/json"
	"testing"
)

// TestExampleSubscriptions tests whether the API documentation's example subscriptions unmarshal correctly to the Subscriptions type
func TestExampleSubscriptions(t *testing.T) {
	subscriptions := &Subscriptions{}
	if err := json.Unmarshal(ExampleSubscriptions, subscriptions); err != nil {
		t.Fatal(err)
	}
}

// TestExampleSubscription tests whether the API documentation's example subscription unmarshals correctly
func TestExampleSubscription(t *testing.T) {
	subscription := &Subscription{}
	if err := json.Unmarshal(ExampleSubscription, subscription); err != nil {
		t.Fatal(err)
	}
}
