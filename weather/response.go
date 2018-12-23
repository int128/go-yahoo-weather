package weather

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
)

// Response represents a response from Weather API.
type Response struct {
	Body    ResponseBody
	Expires time.Time // expires header
}

// ResponseBody represents body of Response.
type ResponseBody struct {
	ResultInfo struct {
		Count       int     `json:"Count"`
		Total       int     `json:"Total"`
		Start       int     `json:"Start"`
		Status      int     `json:"Status"`
		Latency     float64 `json:"Latency"`
		Description string  `json:"Description"`
		Copyright   string  `json:"Copyright"`
	} `json:"ResultInfo"`
	Feature []struct {
		ID       string `json:"Id"`
		Name     string `json:"Name"`
		Geometry struct {
			Type        string            `json:"Type"`
			Coordinates CoordinatesString `json:"Coordinates"`
		} `json:"Geometry"`
		Property struct {
			WeatherAreaCode int `json:"WeatherAreaCode"`
			WeatherList     struct {
				Weather []struct {
					Type     string     `json:"Type"`
					Date     DateString `json:"Date"`
					Rainfall float64    `json:"Rainfall"`
				} `json:"Weather"`
			} `json:"WeatherList"`
		} `json:"Property"`
	} `json:"Feature"`
}

// CoordinatesString represents a coordinates in API specific format.
type CoordinatesString string

// Parse returns a coordinates corresponding to the string.
func (s CoordinatesString) Parse() (Coordinates, error) {
	p := strings.SplitN(string(s), ",", 2)
	if len(p) != 2 {
		return Coordinates{}, errors.Errorf("invalid coordinates string: %s", s)
	}
	lat, lon := p[1], p[0]

	var c Coordinates
	var err error
	c.Latitude, err = strconv.ParseFloat(lat, 64)
	if err != nil {
		return Coordinates{}, errors.Wrapf(err, "error while parsing latitude %s", lat)
	}
	c.Longitude, err = strconv.ParseFloat(lon, 64)
	if err != nil {
		return Coordinates{}, errors.Wrapf(err, "error while parsing longitude %s", lon)
	}
	return c, nil
}

// DateString represents a time in API specific format.
type DateString string

// Parse returns a time.Time corresponding to the string.
func (t DateString) Parse() (time.Time, error) {
	return time.ParseInLocation("200601021504", string(t), Timezone)
}
