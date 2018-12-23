// Package weather provides a client for YOLP Weather API,
// described as https://developer.yahoo.co.jp/webapi/map/openlocalplatform/v1/weather.html
//
package weather

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

var Endpoint = "https://map.yahooapis.jp/weather/V1/place"

// NewClient returns a Client.
func NewClient(clientID string) *Client {
	return &Client{ClientID: clientID}
}

// Client provides access to Weather API.
type Client struct {
	ClientID string       // API Client ID (required)
	Client   *http.Client // Default to http.DefaultClient
	Endpoint string       // Default to Endpoint
}

// Get sends the request and returns the response.
func (c *Client) Get(req *Request) (*Response, error) {
	endpoint := c.Endpoint
	if endpoint == "" {
		endpoint = Endpoint
	}

	q := req.QueryString()
	hreq, err := http.NewRequest("GET", endpoint+"?"+q, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error while creating a HTTP request")
	}
	hreq.Header.Set("user-agent", "Yahoo AppID: "+c.ClientID)

	client := c.Client
	if client == nil {
		client = http.DefaultClient
	}
	hresp, err := client.Do(hreq)
	if err != nil {
		return nil, errors.Wrapf(err, "error while sending a HTTP request")
	}
	defer hresp.Body.Close()
	if hresp.StatusCode != 200 {
		return nil, errors.Errorf("server returned status code %d", hresp.StatusCode)
	}

	var resp Response
	expires, err := http.ParseTime(hresp.Header.Get("expires"))
	if err == nil {
		resp.Expires = expires
	}
	d := json.NewDecoder(hresp.Body)
	if err := d.Decode(&resp.Body); err != nil {
		return nil, errors.Wrapf(err, "error while decoding JSON")
	}
	return &resp, nil
}
