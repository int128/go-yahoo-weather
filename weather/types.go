package weather

import "time"

// Coordinates represents a coordinates in WGS84.
type Coordinates struct {
	Latitude  float64
	Longitude float64
}

// Timezone represents the timezone used in Weather API.
var Timezone = time.FixedZone("Asia/Tokyo", 9*60*60)
