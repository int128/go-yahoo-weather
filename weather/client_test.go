package weather_test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/int128/go-yahoo-weather/weather"
	"github.com/int128/go-yahoo-weather/weather/testdata"
)

var expires = time.Date(2018, 12, 12, 13, 59, 0, 0, time.UTC)

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

		w.Header().Set("content-type", "application/json; charset=UTF-8")
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

func ExampleGetErrorResponse() {
	c := weather.NewClient(os.Getenv("YAHOO_CLIENT_ID"))
	resp, err := c.Get(&weather.Request{})
	if err != nil {
		if errResp := weather.GetErrorResponse(err); errResp != nil {
			if errResp.Code() >= 500 {
				// you can retry here
			}
		}
		log.Fatalf("Could not get weather: %s", err)
	}
	log.Printf("Weather response: %+v", resp)
}

func TestClient_Get_Error400(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("expires", expires.Format(http.TimeFormat))
		w.Header().Set("content-type", "application/json; charset=UTF-8")
		w.WriteHeader(200)
		if _, err := w.Write([]byte(testdata.ErrorResponse400JSON)); err != nil {
			t.Errorf("error while writing body: %s", err)
		}
	}))
	defer s.Close()

	c := weather.NewClient("YAHOO_CLIENT_ID")
	c.Endpoint = s.URL
	resp, err := c.Get(&weather.Request{})
	if resp != nil {
		t.Errorf("resp wants nil but %+v", resp)
	}
	if err == nil {
		t.Fatalf("error wants non-nil but nil")
	}
	if !strings.Contains(err.Error(), "400") {
		t.Errorf("error should contain code")
	}
	t.Log(err)
	errResp := weather.GetErrorResponse(err)
	if errResp == nil {
		t.Errorf("errResp wants non-nil but nil")
	}
	if errResp.Code() != 400 {
		t.Errorf("Code wants 400 but %d", errResp.Code())
	}
	if want := "Bad Request."; errResp.Message() != want {
		t.Errorf("Message wants %s but %s", want, errResp.Message())
	}
}

func TestClient_Get_Error401(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("expires", expires.Format(http.TimeFormat))
		w.Header().Set("content-type", "application/xml")
		w.WriteHeader(401)
		if _, err := w.Write([]byte(testdata.ErrorResponse401XML)); err != nil {
			t.Errorf("error while writing body: %s", err)
		}
	}))
	defer s.Close()

	c := weather.NewClient("YAHOO_CLIENT_ID")
	c.Endpoint = s.URL
	resp, err := c.Get(&weather.Request{})
	if resp != nil {
		t.Errorf("resp wants nil but %+v", resp)
	}
	if err == nil {
		t.Fatalf("error wants non-nil but nil")
	}
	if !strings.Contains(err.Error(), "401") {
		t.Errorf("error should contain code")
	}
	t.Log(err)
	errResp := weather.GetErrorResponse(err)
	if errResp == nil {
		t.Errorf("errResp wants non-nil but nil")
	}
	if errResp.Code() != 401 {
		t.Errorf("Code wants 401 but %d", errResp.Code())
	}
	if want := "Bad Request: Authentication parameters in your request incompleted."; errResp.Message() != want {
		t.Errorf("Message wants %s but %s", want, errResp.Message())
	}
}

func TestClient_Get_Error403(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("expires", expires.Format(http.TimeFormat))
		w.Header().Set("content-type", "application/xml")
		w.WriteHeader(403)
		if _, err := w.Write([]byte(testdata.ErrorResponse403XML)); err != nil {
			t.Errorf("error while writing body: %s", err)
		}
	}))
	defer s.Close()

	c := weather.NewClient("YAHOO_CLIENT_ID")
	c.Endpoint = s.URL
	resp, err := c.Get(&weather.Request{})
	if resp != nil {
		t.Errorf("resp wants nil but %+v", resp)
	}
	if err == nil {
		t.Fatalf("error wants non-nil but nil")
	}
	if !strings.Contains(err.Error(), "403") {
		t.Errorf("error should contain code")
	}
	t.Log(err)
	errResp := weather.GetErrorResponse(err)
	if errResp == nil {
		t.Errorf("errResp wants non-nil but nil")
	}
	if errResp.Code() != 403 {
		t.Errorf("Code wants 403 but %d", errResp.Code())
	}
	if want := "Your Request was Forbidden"; errResp.Message() != want {
		t.Errorf("Message wants %s but %s", want, errResp.Message())
	}
}

func TestGetErrorResponse_NonErrorResponse(t *testing.T) {
	err := fmt.Errorf("some error")
	resp := weather.GetErrorResponse(err)
	if resp != nil {
		t.Errorf("GetErrorResponse wants nil but got %+v", resp)
	}
}
