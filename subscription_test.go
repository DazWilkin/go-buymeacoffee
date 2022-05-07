package client

import (
	"encoding/json"
	"testing"
)

func TestExampleSubscriptions(t *testing.T) {
	subscriptions := &Subscriptions{}
	if err := json.Unmarshal(exampleSubscriptions, subscriptions); err != nil {
		t.Fatal(err)
	}
}
