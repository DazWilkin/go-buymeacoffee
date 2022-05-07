package main

import (
	"log"
	"os"

	client "github.com/DazWilkin/go-buymeacoffee"
)

func main() {

	token := os.Getenv("TOKEN")

	c := client.New(token)
	{
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
	{
		supporter, err := c.Supporter(1652236)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%+v", supporter.ID)
	}
	{
		subscriptions, err := c.Subscriptions(client.All)
		if err != nil {
			log.Fatal(err)
		}

		if subscriptions.Data != nil {
			for _, subscription := range subscriptions.Data {
				log.Printf("%+v", subscription.ID)
			}
		}
	}
}
