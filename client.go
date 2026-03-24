package hisend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Config struct {
	APIKey string
}

type Client struct {
	apiKey  string
	BaseURL string

	Emails  *EmailsService
	Domains *DomainsService
	Routing *RoutingService
	Threads *ThreadsService

	HTTPClient *http.Client
}

func NewClient(config Config) *Client {
	c := &Client{
		apiKey:     config.APIKey,
		BaseURL:    "https://api.hisend.app/v1",
		HTTPClient: &http.Client{},
	}

	c.Emails = &EmailsService{client: c}
	c.Domains = &DomainsService{client: c}
	c.Routing = &RoutingService{client: c}
	c.Threads = &ThreadsService{client: c}

	return c
}

func (c *Client) request(method, endpoint string, body interface{}, v interface{}) error {
	url := fmt.Sprintf("%s%s", c.BaseURL, endpoint)

	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return err
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		resBody, _ := io.ReadAll(res.Body)
		return fmt.Errorf("API request failed with status %d: %s", res.StatusCode, string(resBody))
	}

	if v != nil {
		if err := json.NewDecoder(res.Body).Decode(v); err != nil {
			return err
		}
	}

	return nil
}
