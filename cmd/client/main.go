package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/DazWilkin/go-buymeacoffee/client"
	"github.com/DazWilkin/go-buymeacoffee/text"
	"github.com/DazWilkin/go-buymeacoffee/types"
)

var (
	token = flag.String("token", "", "Developer token")
)

func ID(env string) (uint, error) {
	ID, err := strconv.ParseUint(os.Getenv(env), 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(ID), nil
}
func main() {
	flag.Parse()

	if *token == "" {
		log.Fatal("Flag `--token` is required")
	}

	c := client.New(*token)

	// Supports
	{
		supporters, err := c.Supporters()
		if err != nil {
			log.Fatal(err)
		}

		if len(supporters) != 0 {
			fmt.Println(text.Supporters(supporters))
		}
	}

	// Supporter requires `SUPPORTER`
	{
		ID, err := ID("SUPPORTER")
		if err != nil {
			log.Fatal(err)
		}

		supporter, err := c.Supporter(ID)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%+v\n", supporter.ID)
	}
	// Subscriptions
	{
		subscriptions, err := c.Subscriptions(types.All)
		if err != nil {
			log.Fatal(err)
		}

		if len(subscriptions) != 0 {
			fmt.Println(text.Subscriptions(subscriptions))
		}
	}

	// Subscription requires `SUBSCRIPTION`
	{
		ID, err := ID("SUBSCRIPTION")
		if err != nil {
			log.Fatal(err)
		}

		subscription, err := c.Subscription(ID)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%+v\n", subscription.ID)
	}

	// Purchases
	{
		purchases, err := c.Purchases()
		if err != nil {
			log.Fatal(err)
		}

		if len(purchases) != 0 {
			fmt.Println(text.Purchases(purchases))
		}
	}

	// Purchase requires `PURCHASE`
	{
		ID, err := ID("PURCHASER")
		if err != nil {
			log.Fatal(err)
		}

		purchase, err := c.Purchase(ID)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%+v\n", purchase.ID)
	}
}
