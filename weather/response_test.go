package weather

import (
	"testing"
	"time"
)

func TestCoordinatesString_Parse(t *testing.T) {
	s := CoordinatesString("139.73229,35.663613")
	c, err := s.Parse()
	if err != nil {
		t.Fatalf("CoordinatesString.Parse returned error: %s", err)
	}
	if want := 35.663613; c.Latitude != want {
		t.Errorf("Latitude wants %f but %f", want, c.Latitude)
	}
	if want := 139.73229; c.Longitude != want {
		t.Errorf("Longitude wants %f but %f", want, c.Longitude)
	}
}

func TestDateString_Parse(t *testing.T) {
	s := DateString("201210191610")
	d, err := s.Parse()
	if err != nil {
		t.Fatalf("CoordinatesString.Parse returned error: %s", err)
	}
	if want := 2012; d.Year() != want {
		t.Errorf("Year wants %d but %d", want, d.Year())
	}
	if want := time.October; d.Month() != want {
		t.Errorf("Month wants %d but %d", want, d.Month())
	}
	if want := 19; d.Day() != want {
		t.Errorf("Day wants %d but %d", want, d.Day())
	}
	if want := 16; d.Hour() != want {
		t.Errorf("Hour wants %d but %d", want, d.Hour())
	}
	if want := 10; d.Minute() != want {
		t.Errorf("Minute wants %d but %d", want, d.Minute())
	}
	if _, offset := d.Zone(); offset != 9*60*60 {
		t.Errorf("Zone offset wants %d but %d", 9*60*60, offset)
	}
}
