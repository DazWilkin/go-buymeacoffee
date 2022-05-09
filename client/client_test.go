package client

import (
	"log"
	"testing"

	"github.com/DazWilkin/go-buymeacoffee/types"
)

const (
	testEndpoint        = "http://0.0.0.0:8080"
	testToken    string = ""
)

func ID(env string) uint {
	switch env {
	case "PURCHASE":
		return 30
	case "SUBSCRIPTION":
		return 7979
	case "SUPPORTER":
		return 245731
	default:
		return 0
	}
}

// TestClientPurchases tests the client's Purchases method
func TestClientPurchases(t *testing.T) {
	c := New(testToken)
	c.Endpoint = testEndpoint

	purchases, err := c.Purchases()
	if err != nil {
		t.Fatal(err)
	}

	for _, purchase := range purchases {
		log.Printf("%+v", purchase.ID)
	}
}

// TestClientPurchase tests the client's Purchase method
func TestClientPurchase(t *testing.T) {
	c := New(testToken)
	c.Endpoint = testEndpoint

	ID := ID("PURCHASE")

	purchase, err := c.Purchase(ID)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("%+v", purchase.ID)
}

// TestClientSupporters tests the client's Supporters method
func TestClientSupporters(t *testing.T) {
	c := New(testToken)
	c.Endpoint = testEndpoint

	supporters, err := c.Supporters()
	if err != nil {
		t.Fatal(err)
	}

	for _, supporter := range supporters {
		log.Printf("%+v", supporter.ID)
	}
}

// TestClientSupporter tests the client's Supporter method
func TestClientSupporter(t *testing.T) {
	c := New(testToken)
	c.Endpoint = testEndpoint

	ID := ID("SUPPORTER")

	supporter, err := c.Supporter(ID)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("%+v", supporter.ID)
}

// TestClientSubscriptions tests the client's Subscriptions method
func TestClientSubscriptions(t *testing.T) {
	c := New(testToken)
	c.Endpoint = testEndpoint

	subscriptions, err := c.Subscriptions(types.All)
	if err != nil {
		t.Fatal(err)
	}

	for _, subscription := range subscriptions {
		log.Printf("%+v", subscription.ID)
	}
}

// TestClientSubscription tests the client's Subscription method
func TestClientSubscription(t *testing.T) {
	c := New(testToken)
	c.Endpoint = testEndpoint

	ID := ID("SUBSCRIPTION")

	subscription, err := c.Subscription(ID)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("%+v", subscription.ID)
}
