package weather_test

import (
	"github.com/int128/go-yahoo-weather/weather"
	"github.com/int128/go-yahoo-weather/weather/testdata"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func ExampleClient_Get() {
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

func TestClient_Get(t *testing.T) {
	expires := time.Date(2018, 12, 12, 13, 59, 0, 0, time.UTC)
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Method wants GET but %s", r.Method)
		}
		if r.URL.Path != "/" {
			t.Errorf("Path wants / but %s", r.URL.Path)
		}
		userAgent := r.Header.Get("user-agent")
		if want := "Yahoo AppID: YAHOO_CLIENT_ID"; userAgent != want {
			t.Errorf("user-agent wants %s but %s", want, userAgent)
		}
		coordinates := r.URL.Query().Get("coordinates")
		if want := "139.732293,35.663613 140.728924,41.768674"; coordinates != want {
			t.Errorf("coordinates wants %s but %s", want, coordinates)
		}

		w.Header().Set("expires", expires.Format(http.TimeFormat))
		if _, err := w.Write([]byte(testdata.WeatherResponseJSON)); err != nil {
			t.Errorf("error while writing body: %s", err)
		}
	}))
	defer s.Close()

	c := weather.NewClient("YAHOO_CLIENT_ID")
	c.Endpoint = s.URL
	resp, err := c.Get(&weather.Request{
		Coordinates: []weather.Coordinates{
			{Latitude: 35.663613, Longitude: 139.732293},
			{Latitude: 41.768674, Longitude: 140.728924},
		},
	})
	if err != nil {
		t.Fatalf("Get returned error: %s", err)
	}
	if !resp.Expires.Equal(expires) {
		t.Errorf("Expires wants %v but %v", expires, resp.Expires)
	}
	if len(resp.Body.Feature) != 2 {
		t.Errorf("Feature size wants 2 but %d", len(resp.Body.Feature))
	}
}
