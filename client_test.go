package client

import (
	"log"
	"os"
	"testing"
)

func TestClientSupporters(t *testing.T) {
	token := os.Getenv("TOKEN")
	c := New(token)

	supporters, err := c.Supporters()
	if err != nil {
		log.Fatal(err)
	}

	if supporters.Data != nil {
		for _, supporter := range supporters.Data {
			log.Printf("%+v", supporter.ID)
		}
	}

}
func TestClientSupporter(t *testing.T) {
	token := os.Getenv("TOKEN")
	c := New(token)

	supporter, err := c.Supporter(1652236)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", supporter.ID)
}
func TestClientSubscriptions(t *testing.T) {
	token := os.Getenv("TOKEN")
	c := New(token)

	subscriptions, err := c.Subscriptions(All)
	if err != nil {
		log.Fatal(err)
	}

	if subscriptions.Data != nil {
		for _, subscription := range subscriptions.Data {
			log.Printf("%+v", subscription.ID)
		}
	}
}
