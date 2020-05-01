// Package weather provides a client for YOLP Weather API,
// described as https://developer.yahoo.co.jp/webapi/map/openlocalplatform/v1/weather.html
//
package weather

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
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
		return nil, fmt.Errorf("error while creating a HTTP request: %w", err)
	}
	hreq.Header.Set("user-agent", "Yahoo AppID: "+c.ClientID)

	client := c.Client
	if client == nil {
		client = http.DefaultClient
	}
	hresp, err := client.Do(hreq)
	if err != nil {
		return nil, fmt.Errorf("error while sending a HTTP request: %w", err)
	}
	defer hresp.Body.Close()

	contentType := hresp.Header.Get("content-type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return nil, fmt.Errorf("invalid content-type header: %w", err)
	}
	switch mediaType {
	case "application/json":
		var resp Response
		expires, err := http.ParseTime(hresp.Header.Get("expires"))
		if err == nil {
			resp.Expires = expires
		}
		d := json.NewDecoder(hresp.Body)
		if err := d.Decode(&resp.Body); err != nil {
			return nil, fmt.Errorf("error while decoding JSON response: %w", err)
		}
		if errResp := resp.Body.Error; errResp.Code() >= 300 || hresp.StatusCode >= 300 {
			if errResp.Code() == 0 {
				errResp.CodeValue = hresp.StatusCode
			}
			return nil, fmt.Errorf("%w", &resp.Body.Error)
		}
		return &resp, nil

	case "application/xml":
		var errResp errorResponse
		d := xml.NewDecoder(hresp.Body)
		if err := d.Decode(&errResp); err != nil {
			return nil, fmt.Errorf("error while decoding XML response: %w", err)
		}
		if errResp.Code() == 0 {
			errResp.CodeValue = hresp.StatusCode
		}
		return nil, &errResp

	default:
		b, _ := ioutil.ReadAll(hresp.Body)
		return nil, fmt.Errorf("received unknown content-type %s (body=%s)", mediaType, string(b))
	}
}
