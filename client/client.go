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

// Subscriptions is a method that returns a list of subscriptions
func (c *Client) Subscriptions(status types.Status) (*types.Subscriptions, error) {
	// Add Status to QueryString
	v := xurl.Values{}
	v.Add("status", status.String())
	querystring := v.Encode()

	url := fmt.Sprintf("%s/subscriptions?%s", base, querystring)

	j, err := c.Do(url)
	if err != nil {
		return &types.Subscriptions{}, err
	}

	subscriptions := &types.Subscriptions{}
	if err := json.Unmarshal(j, subscriptions); err != nil {
		return &types.Subscriptions{}, err
	}

	return subscriptions, nil
}

// Supporter is a method that given a supporter ID returns a supporter
func (c *Client) Supporter(ID uint) (*types.Supporter, error) {
	url := fmt.Sprintf("%s/supporters/%d", base, ID)

	j, err := c.Do(url)
	if err != nil {
		log.Fatal(err)
	}

	supporter := &types.Supporter{}
	if err := json.Unmarshal(j, supporter); err != nil {
		return &types.Supporter{}, err
	}

	return supporter, nil
}

// Supporters is a method that returns a list of supporters
func (c *Client) Supporters() (*types.Supporters, error) {
	url := fmt.Sprintf("%s/supporters", base)

	j, err := c.Do(url)
	if err != nil {
		log.Fatal(err)
	}

	supporters := &types.Supporters{}
	if err := json.Unmarshal(j, supporters); err != nil {
		return &types.Supporters{}, err
	}

	return supporters, nil
}