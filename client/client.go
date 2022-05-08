package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	xurl "net/url"

	"github.com/DazWilkin/go-buymeacoffee/types"
)

const (
	base string = "https://developers.buymeacoffee.com/api/v1"
)

// Client is a type that represents an HTTP client of Buy Me A Coffee API
type Client struct {
	Token  string
	client *http.Client
}

// New is a function that creates a new Client
func New(token string) Client {
	return Client{
		Token:  token,
		client: &http.Client{},
	}
}

// Do is a method that wraps the http.Do method for this client
func (c *Client) Do(url string) ([]byte, error) {
	rqst, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Add Authorization
	rqst.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	resp, err := c.client.Do(rqst)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

// Subscription is a method that given a subscription ID returns a subscription
func (c *Client) Subscription(ID uint) (types.Subscription, error) {
	url := fmt.Sprintf("%s/subscription/%d", base, ID)

	j, err := c.Do(url)
	if err != nil {
		return types.Subscription{}, err
	}

	subscription := types.Subscription{}
	if err := json.Unmarshal(j, &subscription); err != nil {
		return types.Subscription{}, err
	}

	return subscription, nil
}

// Subscriptions is a method that returns a list of subscriptions
func (c *Client) Subscriptions(status types.Status) ([]types.Subscription, error) {
	// Add Status to QueryString
	v := xurl.Values{}
	v.Add("status", status.String())
	querystring := v.Encode()

	url := fmt.Sprintf("%s/subscriptions?%s", base, querystring)

	ss := []types.Subscription{}

	// Loop at least once
	// Until NextPageURL is null
	for {
		j, err := c.Do(url)
		if err != nil {
			return ss, err
		}

		subscriptions := &types.Subscriptions{}
		if err := json.Unmarshal(j, subscriptions); err != nil {
			return ss, err
		}

		ss = append(ss, subscriptions.Data...)

		if subscriptions.NextPageURL == "" {
			break
		}
	}

	return ss, nil
}

// Supporter is a method that given a supporter ID returns a supporter
func (c *Client) Supporter(ID uint) (types.Supporter, error) {
	url := fmt.Sprintf("%s/supporters/%d", base, ID)

	j, err := c.Do(url)
	if err != nil {
		log.Fatal(err)
	}

	supporter := types.Supporter{}
	if err := json.Unmarshal(j, &supporter); err != nil {
		return types.Supporter{}, err
	}

	return supporter, nil
}

// Supporters is a method that returns a list of supporters
func (c *Client) Supporters() ([]types.Supporter, error) {
	url := fmt.Sprintf("%s/supporters", base)

	// Return value
	ss := []types.Supporter{}

	// Loop at least once
	// Until NextPageURL is null
	for {
		j, err := c.Do(url)
		if err != nil {
			log.Fatal(err)
		}

		supporters := &types.Supporters{}
		if err := json.Unmarshal(j, supporters); err != nil {
			return ss, err
		}

		ss = append(ss, supporters.Data...)

		if supporters.NextPageURL == "" {
			break
		}
	}

	return ss, nil
}
