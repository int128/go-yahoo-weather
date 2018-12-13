package weather_test

import (
	"encoding/json"
	"github.com/go-test/deep"
	"github.com/int128/go-yahoo-weather/weather"
	"github.com/int128/go-yahoo-weather/weather/testdata"
	"log"
	"os"
	"testing"
	"time"
)

func ExampleParse() {
	c := weather.NewClient(os.Getenv("YAHOO_CLIENT_ID"))
	resp, err := c.Get(&weather.Request{
		Coordinates: []weather.Coordinates{
			{Latitude: 35.663613, Longitude: 139.732293},
		},
	})
	if err != nil {
		log.Fatalf("Error while getting weather: %s", err)
	}
	weathers, err := weather.Parse(resp)
	if err != nil {
		log.Fatalf("Error while parsing weather response: %s", err)
	}
	log.Printf("Weathers: %+v", weathers)
}

func TestParse(t *testing.T) {
	var resp weather.Response
	if err := json.Unmarshal([]byte(testdata.WeatherResponseJSON), &resp.Body); err != nil {
		t.Fatalf("error while decoding JSON: %s", err)
	}
	weathers, err := weather.Parse(&resp)
	if err != nil {
		t.Fatalf("Parse returned error: %s", err)
	}
	want := []weather.Weather{
		{
			Coordinates: weather.Coordinates{Latitude: 35.663613, Longitude: 139.73229},
			Events: []weather.Event{
				{Time: time.Date(2018, 12, 12, 13, 5, 0, 0, weather.Timezone)},
				{Time: time.Date(2018, 12, 12, 13, 15, 0, 0, weather.Timezone), Forecast: true},
				{Time: time.Date(2018, 12, 12, 13, 25, 0, 0, weather.Timezone), Forecast: true},
				{Time: time.Date(2018, 12, 12, 13, 35, 0, 0, weather.Timezone), Forecast: true},
				{Time: time.Date(2018, 12, 12, 13, 45, 0, 0, weather.Timezone), Forecast: true},
				{Time: time.Date(2018, 12, 12, 13, 55, 0, 0, weather.Timezone), Forecast: true},
				{Time: time.Date(2018, 12, 12, 14, 5, 0, 0, weather.Timezone), Forecast: true},
			},
		}, {
			Coordinates: weather.Coordinates{Latitude: 41.768674, Longitude: 140.72892},
			Events: []weather.Event{
				{Time: time.Date(2018, 12, 12, 13, 5, 0, 0, weather.Timezone), Rainfall: 0.35},
				{Time: time.Date(2018, 12, 12, 13, 15, 0, 0, weather.Timezone), Forecast: true, Rainfall: 0.45},
				{Time: time.Date(2018, 12, 12, 13, 25, 0, 0, weather.Timezone), Forecast: true, Rainfall: 1.15},
				{Time: time.Date(2018, 12, 12, 13, 35, 0, 0, weather.Timezone), Forecast: true, Rainfall: 0.45},
				{Time: time.Date(2018, 12, 12, 13, 45, 0, 0, weather.Timezone), Forecast: true, Rainfall: 1.85},
				{Time: time.Date(2018, 12, 12, 13, 55, 0, 0, weather.Timezone), Forecast: true},
				{Time: time.Date(2018, 12, 12, 14, 5, 0, 0, weather.Timezone), Forecast: true},
			},
		},
	}
	if diff := deep.Equal(want, weathers); diff != nil {
		t.Error(diff)
	}
}
