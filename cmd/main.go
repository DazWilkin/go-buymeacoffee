package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/DazWilkin/go-buymeacoffee/client"
	"github.com/DazWilkin/go-buymeacoffee/types"
)

func main() {

	token := os.Getenv("TOKEN")
	supporterID, err := strconv.ParseUint(os.Getenv("SUPPORTER"), 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	c := client.New(token)

	{
		supporters, err := c.Supporters()
		if err != nil {
			log.Fatal(err)
		}

		if len(supporters) != 0 {
			fmt.Println(types.SupportersToText(supporters))
		}
	}
	{
		supporter, err := c.Supporter(uint(supporterID))
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%+v", supporter.ID)
	}
	{
		subscriptions, err := c.Subscriptions(types.All)
		if err != nil {
			log.Fatal(err)
		}

		if len(subscriptions) != 0 {
			fmt.Println(types.SubscriptionsToText(subscriptions))
		}
	}
}
