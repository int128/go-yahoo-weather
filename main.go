package main

import (
	"github.com/int128/go-yahoo-weather/weather"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

type loggingTransport struct{}

func (t *loggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	dump, err := httputil.DumpRequestOut(req, true)
	if err == nil {
		log.Printf("REQUEST:\n%s", string(dump))
	}
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return resp, err
	}
	dump, err = httputil.DumpResponse(resp, true)
	if err == nil {
		log.Printf("RESPONSE:\n%s", string(dump))
	}
	return resp, nil
}

func main() {
	hc := http.Client{Transport: &loggingTransport{}}
	c := weather.NewClient(os.Getenv("YAHOO_CLIENT_ID"))
	c.Client = &hc

	req := weather.Request{
		Coordinates: []weather.Coordinates{
			{Latitude: 35.663613, Longitude: 139.732293},  // Roppongi
			{Latitude: 41.7686738, Longitude: 140.728924}, // Hakodate
		},
	}
	log.Printf("Request: %+v", req)
	resp, err := c.Get(&req)
	if err != nil {
		log.Fatalf("Could not get weather: %s", err)
	}
	log.Printf("Response: %+v", resp)
}
