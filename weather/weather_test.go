package weather_test

import (
	"encoding/json"
	"github.com/int128/go-yahoo-weather/weather"
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
	if err := json.Unmarshal([]byte(weatherResponseJSON), &resp.Body); err != nil {
		t.Fatalf("error while decoding JSON: %s", err)
	}

	weathers, err := weather.Parse(&resp)
	if err != nil {
		t.Fatalf("Parse returned error: %s", err)
	}
	if len(weathers) != 2 {
		t.Errorf("len wants 2 but %d", len(weathers))
	}

	w := weathers[0]
	if want := 35.663613; w.Coordinates.Latitude != want {
		t.Errorf("Latitude wants %f but %f", want, w.Coordinates.Latitude)
	}
	if want := 139.73229; w.Coordinates.Longitude != want {
		t.Errorf("Longitude wants %f but %f", want, w.Coordinates.Longitude)
	}
	if len(w.Events) != 7 {
		t.Errorf("len(Events) wants 7 but %d", len(w.Events))
	}

	e := w.Events[6]
	if want := time.Date(2018, 12, 12, 14, 5, 0, 0, weather.Timezone); !e.Time.Equal(want) {
		t.Errorf("Time wants %s but %s", want, e.Time)
	}
	if !e.Forecast {
		t.Errorf("Forecast wants true but false")
	}
}
