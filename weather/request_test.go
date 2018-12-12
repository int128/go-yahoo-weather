package weather

import (
	"testing"
	"time"
)

func TestRequest_Values_Zero(t *testing.T) {
	r := Request{}
	s := r.QueryString()
	w := "coordinates=&output=json"
	if w != s {
		t.Errorf("Values wants %s but %s", w, s)
	}
}

func TestRequest_Values_Coordinates(t *testing.T) {
	r := Request{
		Coordinates: []Coordinates{{
			Latitude:  35.663613,
			Longitude: 139.732293,
		}},
	}
	s := r.QueryString()
	w := "coordinates=139.732293%2C35.663613&output=json"
	if w != s {
		t.Errorf("Values wants %s but %s", w, s)
	}
}

func TestRequest_Values_MultiCoordinates(t *testing.T) {
	r := Request{
		Coordinates: []Coordinates{
			{
				Latitude:  35.663613,
				Longitude: 139.732293,
			},
			{
				Latitude:  35.681167,
				Longitude: 139.767052,
			},
		},
	}
	s := r.QueryString()
	w := "coordinates=139.732293%2C35.663613%20139.767052%2C35.681167&output=json"
	if w != s {
		t.Errorf("Values wants %s but %s", w, s)
	}
}

func TestRequest_Values_DateTime_JST(t *testing.T) {
	r := Request{
		DateTime: time.Date(2018, 11, 30, 12, 34, 56, 0, Timezone),
	}
	s := r.QueryString()
	w := "coordinates=&date=201811301234&output=json"
	if w != s {
		t.Errorf("Values wants %s but %s", w, s)
	}
}

func TestRequest_Values_DateTime_UTC(t *testing.T) {
	r := Request{
		DateTime: time.Date(2018, 11, 30, 12, 34, 56, 0, time.UTC),
	}
	s := r.QueryString()
	w := "coordinates=&date=201811302134&output=json"
	if w != s {
		t.Errorf("Values wants %s but %s", w, s)
	}
}

func TestRequest_Values_PastHours(t *testing.T) {
	r := Request{
		PastHours: 1,
	}
	s := r.QueryString()
	w := "coordinates=&output=json&past=1"
	if w != s {
		t.Errorf("Values wants %s but %s", w, s)
	}
}

func TestRequest_Values_IntervalMinutes(t *testing.T) {
	r := Request{
		IntervalMinutes: 5,
	}
	s := r.QueryString()
	w := "coordinates=&interval=5&output=json"
	if w != s {
		t.Errorf("Values wants %s but %s", w, s)
	}
}
