package client

import (
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/DazWilkin/go-buymeacoffee/types"
)

// TestClientSupporters tests the client's Supporters method
func TestClientSupporters(t *testing.T) {
	token := os.Getenv("TOKEN")
	c := New(token)

	supporters, err := c.Supporters()
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	c := New(token)

	supporter, err := c.Supporter(uint(supporterID))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", supporter.ID)
}

// TestClientSubscriptions tests the client's Subscriptions method
func TestClientSubscriptions(t *testing.T) {
	token := os.Getenv("TOKEN")
	c := New(token)

	subscriptions, err := c.Subscriptions(types.All)
	if err != nil {
		log.Fatal(err)
	}

	for _, subscription := range subscriptions {
		log.Printf("%+v", subscription.ID)
	}
}
