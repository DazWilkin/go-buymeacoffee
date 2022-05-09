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
	endpoint string = "https://developers.buymeacoffee.com/api/v1"
)

// Client is a type that represents an HTTP client of Buy Me A Coffee API
type Client struct {
	Endpoint string
	Token    string
	client   *http.Client
}

// New is a function that creates a new Client
func New(token string) Client {
	return Client{
		Endpoint: endpoint,
		Token:    token,
		client:   &http.Client{},
	}
}

// Do is a method that wraps the http.Do method for this client
func (c *Client) Do(url string) ([]byte, error) {
	// For testing, the url provided may result from NextPageURL
	// And **not** match the Client's Endpoint
	// Check for this discrepancy
	l := len(c.Endpoint)
	if len(url) >= l {
		if url[0:l] != c.Endpoint {
			log.Printf("Discrepancy between URL (%s) and client's endpoint (%s)", url, c.Endpoint)
		}
	}

	rqst, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}

	// Add Authorization
	rqst.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	resp, err := c.client.Do(rqst)
	if err != nil {
		return []byte{}, err
	}

	return ioutil.ReadAll(resp.Body)
}

// Purchase is a method that given a purchase ID returns a purchase
func (c *Client) Purchase(ID uint) (types.Purchase, error) {
	url := fmt.Sprintf("%s/extras/%d", c.Endpoint, ID)

	purchase := types.Purchase{}

	j, err := c.Do(url)
	if err != nil {
		return purchase, err
	}

	if err := json.Unmarshal(j, &purchase); err != nil {
		return purchase, err
	}

	return purchase, nil

}

// Purchases is a method that returns a list of purchases
func (c *Client) Purchases() ([]types.Purchase, error) {
	url := fmt.Sprintf("%s/extras", c.Endpoint)

	pp := []types.Purchase{}

	// Loop at least once
	// Until NextPageURL is null
	for {
		j, err := c.Do(url)
		if err != nil {
			return pp, err
		}

		purchases := types.Purchases{}
		if err := json.Unmarshal(j, &purchases); err != nil {
			return pp, err
		}

		pp = append(pp, purchases.Data...)

		if purchases.NextURL == "" {
			break
		}

		// Update URL to point to the next page
		url = purchases.NextURL
	}

	return pp, nil
}

// Subscription is a method that given a subscription ID returns a subscription
func (c *Client) Subscription(ID uint) (types.Subscription, error) {
	url := fmt.Sprintf("%s/subscriptions/%d", c.Endpoint, ID)

	subscription := types.Subscription{}

	j, err := c.Do(url)
	if err != nil {
		return subscription, err
	}

	if err := json.Unmarshal(j, &subscription); err != nil {
		return subscription, err
	}

	return subscription, nil
}

// Subscriptions is a method that returns a list of subscriptions
func (c *Client) Subscriptions(status types.Status) ([]types.Subscription, error) {
	// Add Status to QueryString
	v := xurl.Values{}
	v.Add("status", status.String())
	querystring := v.Encode()

	url := fmt.Sprintf("%s/subscriptions?%s", c.Endpoint, querystring)

	ss := []types.Subscription{}

	// Loop at least once
	// Until NextPageURL is null
	for {
		j, err := c.Do(url)
		if err != nil {
			return ss, err
		}

		subscriptions := types.Subscriptions{}
		if err := json.Unmarshal(j, &subscriptions); err != nil {
			return ss, err
		}

		ss = append(ss, subscriptions.Data...)

		if subscriptions.NextURL == "" {
			break
		}

		// Update URL to point to the next page
		url = subscriptions.NextURL

	}

	return ss, nil
}

// Supporter is a method that given a supporter ID returns a supporter
func (c *Client) Supporter(ID uint) (types.Supporter, error) {
	url := fmt.Sprintf("%s/supporters/%d", c.Endpoint, ID)

	supporter := types.Supporter{}

	j, err := c.Do(url)
	if err != nil {
		return supporter, err
	}

	if err := json.Unmarshal(j, &supporter); err != nil {
		return supporter, err
	}

	return supporter, nil
}

// Supporters is a method that returns a list of supporters
func (c *Client) Supporters() ([]types.Supporter, error) {
	url := fmt.Sprintf("%s/supporters", c.Endpoint)

	// Return value
	ss := []types.Supporter{}

	// Loop at least once
	// Until NextPageURL is null
	for {
		j, err := c.Do(url)
		if err != nil {
			return ss, err
		}

		supporters := types.Supporters{}
		if err := json.Unmarshal(j, &supporters); err != nil {
			return ss, err
		}

		ss = append(ss, supporters.Data...)

		if supporters.NextURL == "" {
			break
		}

		// Update URL to point to the next page
		url = supporters.NextURL

	}

	return ss, nil
}
