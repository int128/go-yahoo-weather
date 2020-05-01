package weather

import (
	"fmt"
	"time"
)

// Weather represents weather at the coordinates.
type Weather struct {
	Coordinates Coordinates
	Events      []Event
}

// Event represents an observation or forecast at time.
type Event struct {
	Time     time.Time
	Forecast bool
	Rainfall float64
}

// Coordinates represents a coordinates in WGS84.
type Coordinates struct {
	Latitude  float64
	Longitude float64
}

// Timezone represents the timezone used in Weather API.
var Timezone = time.FixedZone("Asia/Tokyo", 9*60*60)

// Parse returns a list of Weather corresponding to the Response.
func Parse(resp *Response) ([]Weather, error) {
	weathers := make([]Weather, 0)
	for _, respFeature := range resp.Body.Feature {
		var weather Weather
		var err error

		weather.Coordinates, err = respFeature.Geometry.Coordinates.Parse()
		if err != nil {
			return nil, fmt.Errorf("invalid coordinates: %w", err)
		}

		for _, respWeather := range respFeature.Property.WeatherList.Weather {
			var event Event
			event.Rainfall = respWeather.Rainfall
			event.Time, err = respWeather.Date.Parse()
			if err != nil {
				return nil, fmt.Errorf("invalid date: %w", err)
			}
			switch respWeather.Type {
			case "observation":
			case "forecast":
				event.Forecast = true
			default:
				return nil, fmt.Errorf("unknown weather type: %s", respWeather.Type)
			}

			weather.Events = append(weather.Events, event)
		}

		weathers = append(weathers, weather)
	}
	return weathers, nil
}
