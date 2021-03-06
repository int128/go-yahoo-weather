package weather

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

// Request represents a request for Weather API.
type Request struct {
	Coordinates     []Coordinates // list of coordinates, up to 10 (required)
	DateTime        time.Time     // default to current time
	PastHours       int           // 0 (default), 1 or 2
	IntervalMinutes int           // 10 (default) or 5
}

// Values returns query parameters corresponding to the request.
func (r *Request) Values() url.Values {
	v := make(url.Values)
	v.Set("output", "json")
	s := make([]string, len(r.Coordinates))
	for i, c := range r.Coordinates {
		s[i] = fmt.Sprintf("%f,%f", c.Longitude, c.Latitude)
	}
	v.Set("coordinates", strings.Join(s, " "))
	if !r.DateTime.IsZero() {
		s := r.DateTime.In(Timezone).Format("200601021504")
		v.Set("date", s)
	}
	if r.PastHours != 0 {
		v.Set("past", fmt.Sprintf("%d", r.PastHours))
	}
	if r.IntervalMinutes != 0 {
		v.Set("interval", fmt.Sprintf("%d", r.IntervalMinutes))
	}
	return v
}

// QueryString returns the query string.
// Note that spaces need to encoded as %20 instead of +.
func (r *Request) QueryString() string {
	q := r.Values().Encode()
	return strings.Replace(q, "+", "%20", -1)
}
