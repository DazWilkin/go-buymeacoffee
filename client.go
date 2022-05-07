package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	base string = "https://developers.buymeacoffee.com/api/v1"
)

type Client struct {
	Token  string
	client *http.Client
}

func New(token string) Client {
	return Client{
		Token:  token,
		client: &http.Client{},
	}
}
func (c *Client) Do(rqst *http.Request) (*http.Response, error) {

}

func (c *Client) Subscriptions(status Status) (*Subscriptions, error) {
	url := fmt.Sprintf("%s/subscriptions", base)

	rqst, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Add Authorization
	rqst.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))

	// Add Status to QueryString
	querystring := rqst.URL.Query()
	querystring.Add("status", status.String())
	rqst.URL.RawQuery = querystring.Encode()

	resp, err := c.client.Do(rqst)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	subscriptions := &Subscriptions{}
	if err := json.Unmarshal(body, subscriptions); err != nil {
		return &Subscriptions{}, err
	}

	return subscriptions, nil
}
func (c *Client) Supporter(ID uint) (*Supporter, error) {
	url := fmt.Sprintf("%s/supporters/%d", base, ID)

	rqst, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	rqst.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))

	resp, err := c.client.Do(rqst)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	supporter := &Supporter{}
	if err := json.Unmarshal(body, supporter); err != nil {
		return &Supporter{}, err
	}

	return supporter, nil
}
func (c *Client) Supporters() (*Supporters, error) {
	url := fmt.Sprintf("%s/supporters", base)

	rqst, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	rqst.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))

	resp, err := c.client.Do(rqst)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	supportersListResp := &Supporters{}
	if err := json.Unmarshal(body, supportersListResp); err != nil {
		return &Supporters{}, err
	}

	return supportersListResp, nil
}
