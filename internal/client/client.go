// Package client takes care of JSON & XML API requests.
package client

import (
	"encoding/json"
	"net/http"

	"github.com/kinduff/prometheus_exporter_template/config"
	log "github.com/sirupsen/logrus"
)

// Client is a struct that contains an HTTP Client
type Client struct {
	httpClient http.Client
}

// NewClient provides an interface to make HTTP requests to an API.
func NewClient() *Client {
	return &Client{
		httpClient: http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}
}

type DoAPIRequestOptions struct {
	Id string
}

// DoAPIRequest allows to make requests to an API by standarizing how it receives
// parameters, and to which endpoint it should do the call.
func (client *Client) DoAPIRequest(endpoint string, config *config.Config, target interface{}, opts *DoAPIRequestOptions) error {
	req, err := http.NewRequest("GET", getAPIEndpoint(endpoint, opts, config), nil)
	if err != nil {
		log.Fatalf("An error has occurred when creating HTTP request %v", err)

		return err
	}

	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Accept":        []string{"application/json"},
		"Authorization": []string{"Bearer " + config.APIKey},
	}

	req.URL.RawQuery = getAPIQueryParams(endpoint, opts, req)

	log.Infof("Sending HTTP request to %s", req.URL.String())

	resp, err := client.httpClient.Do(req)
	if err != nil || !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		log.Fatalf("An error has occurred during retrieving stats %v", err)

		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func getAPIEndpoint(endpoint string, opts *DoAPIRequestOptions, config *config.Config) string {
	var path string
	baseUrl := config.BaseURL

	switch endpoint {
	case "demo":
		path = "/posts/" + opts.Id
	}

	return baseUrl + path
}

func getAPIQueryParams(endpoint string, opts *DoAPIRequestOptions, req *http.Request) string {
	q := req.URL.Query()

	switch endpoint {
	case "demo":
		q.Add("refresh", "true")
	}

	return q.Encode()
}
