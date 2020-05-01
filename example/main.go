package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/int128/go-yahoo-weather/weather"
)

func main() {
	c := weather.NewClient(os.Getenv("YAHOO_CLIENT_ID"))
	c.Client = &http.Client{Transport: &loggingTransport{}}

	req := weather.Request{
		Coordinates: []weather.Coordinates{
			{Latitude: 35.663613, Longitude: 139.732293},  // Roppongi
			{Latitude: 41.7686738, Longitude: 140.728924}, // Hakodate
		},
	}
	resp, err := c.Get(&req)
	if err != nil {
		log.Fatalf("Could not get weather: %s", err)
	}
	weathers, err := weather.Parse(resp)
	if err != nil {
		log.Fatalf("Could not parse weather response: %s", err)
	}
	for i, w := range weathers {
		fmt.Printf("Weather #%d at (%f,%f)\n", i, w.Coordinates.Latitude, w.Coordinates.Longitude)
		for _, e := range w.Events {
			forecastMark := " "
			if e.Forecast {
				forecastMark = "*"
			}
			fmt.Printf("| %s %s | %5.2f mm/h |\n", e.Time.Format("15:04"), forecastMark, e.Rainfall)
		}
	}
}

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
