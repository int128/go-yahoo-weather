package weather_test

import (
	"github.com/int128/go-yahoo-weather/weather"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func ExampleNewClient() {
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

const weatherResponseJSON = `{"ResultInfo":{"Count":2,"Total":2,"Start":1,"Status":200,"Latency":0.005322,"Description":"","Copyright":"(C) Yahoo Japan Corporation."},"Feature":[{"Id":"201812121305_139.73229_35.663613","Name":"地点(139.73229,35.663613)の2018年12月12日 13時05分から60分間の天気情報","Geometry":{"Type":"point","Coordinates":"139.73229,35.663613"},"Property":{"WeatherAreaCode":4410,"WeatherList":{"Weather":[{"Type":"observation","Date":"201812121305","Rainfall":0.00},{"Type":"forecast","Date":"201812121315","Rainfall":0.00},{"Type":"forecast","Date":"201812121325","Rainfall":0.00},{"Type":"forecast","Date":"201812121335","Rainfall":0.00},{"Type":"forecast","Date":"201812121345","Rainfall":0.00},{"Type":"forecast","Date":"201812121355","Rainfall":0.00},{"Type":"forecast","Date":"201812121405","Rainfall":0.00}]}}},{"Id":"201812121305_140.72892_41.768674","Name":"地点(140.72892,41.768674)の2018年12月12日 13時05分から60分間の天気情報","Geometry":{"Type":"point","Coordinates":"140.72892,41.768674"},"Property":{"WeatherAreaCode":2300,"WeatherList":{"Weather":[{"Type":"observation","Date":"201812121305","Rainfall":0.35},{"Type":"forecast","Date":"201812121315","Rainfall":0.45},{"Type":"forecast","Date":"201812121325","Rainfall":1.15},{"Type":"forecast","Date":"201812121335","Rainfall":0.45},{"Type":"forecast","Date":"201812121345","Rainfall":1.85},{"Type":"forecast","Date":"201812121355","Rainfall":0.00},{"Type":"forecast","Date":"201812121405","Rainfall":0.00}]}}}]}`

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
		if _, err := w.Write([]byte(weatherResponseJSON)); err != nil {
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
