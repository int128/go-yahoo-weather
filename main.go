package main

import (
	"github.com/int128/go-yahoo-weather/weather"
	"log"
	"os"
)

func main() {
	c := weather.NewClient(os.Getenv("YAHOO_CLIENT_ID"))
	resp, err := c.Get(&weather.Request{
		Coordinates: []weather.Coordinates{
			{Latitude: 35.663613, Longitude: 139.732293},
		},
	})
	if err != nil {
		log.Fatalf("Could not get weather: %s", err)
	}
	log.Printf("Weather response: %+v", resp)
}
