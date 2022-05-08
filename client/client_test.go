package client

import (
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/DazWilkin/go-buymeacoffee/types"
)

// TestClientPurchases tests the client's Purchases method
func TestClientPurchases(t *testing.T) {
	token := os.Getenv("TOKEN")
	c := New(token)

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
	token := os.Getenv("TOKEN")
	purchaseID, err := strconv.ParseUint(os.Getenv("PURCHASE"), 10, 32)
	if err != nil {
		t.Fatal(err)
	}

	c := New(token)

	purchase, err := c.Purchase(uint(purchaseID))
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("%+v", purchase.ID)
}

// TestClientSupporters tests the client's Supporters method
func TestClientSupporters(t *testing.T) {
	token := os.Getenv("TOKEN")
	c := New(token)

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
	token := os.Getenv("TOKEN")
	supporterID, err := strconv.ParseUint(os.Getenv("SUPPORTER"), 10, 32)
	if err != nil {
		t.Fatal(err)
	}

	c := New(token)

	supporter, err := c.Supporter(uint(supporterID))
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("%+v", supporter.ID)
}

// TestClientSubscriptions tests the client's Subscriptions method
func TestClientSubscriptions(t *testing.T) {
	token := os.Getenv("TOKEN")
	c := New(token)

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
	token := os.Getenv("TOKEN")
	subscriptionID, err := strconv.ParseUint(os.Getenv("SUBSCRIPTION"), 10, 32)
	if err != nil {
		t.Fatal(err)
	}

	c := New(token)

	subscription, err := c.Subscription(uint(subscriptionID))
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("%+v", subscription.ID)
}
