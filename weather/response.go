package weather

import (
	"errors"
	"fmt"
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
	Error errorResponse `json:"Error"`
}

type errorResponse struct {
	CodeValue    int    `json:"Code" xml:"Code"`
	MessageValue string `json:"Message" xml:"Message"`
}

func (e *errorResponse) Error() string {
	return fmt.Sprintf("error from Weather API: code=%d, message=%s", e.Code(), e.Message())
}

func (e *errorResponse) Code() int {
	return e.CodeValue
}

func (e *errorResponse) Message() string {
	return strings.TrimSpace(e.MessageValue)
}

// ErrorResponse provides details of error response.
type ErrorResponse interface {
	Code() int       // Error code in response body or status code, such as 400, 401 or 403
	Message() string // Error message
}

// GetErrorResponse returns ErrorResponse if API returned an error response.
func GetErrorResponse(err error) ErrorResponse {
	var v ErrorResponse
	if errors.As(err, &v) {
		return v
	}
	return nil
}

// CoordinatesString represents a coordinates in API specific format.
type CoordinatesString string

// Parse returns a coordinates corresponding to the string.
func (s CoordinatesString) Parse() (Coordinates, error) {
	p := strings.SplitN(string(s), ",", 2)
	if len(p) != 2 {
		return Coordinates{}, fmt.Errorf("invalid coordinates string: %s", s)
	}
	lat, lon := p[1], p[0]

	var c Coordinates
	var err error
	c.Latitude, err = strconv.ParseFloat(lat, 64)
	if err != nil {
		return Coordinates{}, fmt.Errorf("error while parsing latitude %s: %w", lat, err)
	}
	c.Longitude, err = strconv.ParseFloat(lon, 64)
	if err != nil {
		return Coordinates{}, fmt.Errorf("error while parsing longitude %s: %w", lon, err)
	}
	return c, nil
}

// DateString represents a time in API specific format.
type DateString string

// Parse returns a time.Time corresponding to the string.
func (t DateString) Parse() (time.Time, error) {
	return time.ParseInLocation("200601021504", string(t), Timezone)
}
