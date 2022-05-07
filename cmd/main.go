package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	client "github.com/DazWilkin/go-buymeacoffee"
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

		if supporters.Data != nil {
			fmt.Println(supporters.Text())
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
		subscriptions, err := c.Subscriptions(client.All)
		if err != nil {
			log.Fatal(err)
		}

		if subscriptions.Data != nil {
			fmt.Println(subscriptions.Text())
		}
	}
}
