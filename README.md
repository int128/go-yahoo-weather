# go-yahoo-weather [![GoDoc](https://godoc.org/github.com/int128/go-yahoo-weather/weather?status.svg)](https://godoc.org/github.com/int128/go-yahoo-weather/weather) [![CircleCI](https://circleci.com/gh/int128/go-yahoo-weather.svg?style=shield)](https://circleci.com/gh/int128/go-yahoo-weather) [![codecov](https://codecov.io/gh/int128/go-yahoo-weather/branch/master/graph/badge.svg)](https://codecov.io/gh/int128/go-yahoo-weather)

This is a Go package for [Yahoo Japan Weather API](https://developer.yahoo.co.jp/webapi/map/openlocalplatform/v1/weather.html).

See [GoDoc](https://godoc.org/github.com/int128/go-yahoo-weather/weather).


## Getting Started

You need to get a Client ID on [Yahoo Japan Developers Network](https://developer.yahoo.co.jp).

You can get weather observation and forecast as follows:

```go
package main

import (
	"log"
	"os"

	"github.com/int128/go-yahoo-weather/weather"
)

func main() {
	c := weather.NewClient(os.Getenv("YAHOO_CLIENT_ID"))
	resp, err := c.Get(&weather.Request{
		Coordinates: []weather.Coordinates{
			{Latitude: 35.663613, Longitude: 139.732293},  // Roppongi
			{Latitude: 41.7686738, Longitude: 140.728924}, // Hakodate
		},
	})
	if err != nil {
		log.Fatalf("Could not get weather: %s", err)
	}
	log.Printf("Weather response: %+v", resp)

	weathers, err := weather.Parse(resp)
	if err != nil {
		log.Fatalf("Could not parse weather response: %s", err)
	}
	log.Printf("Weathers: %+v", weathers)
}
```

See also [example/main.go](example/main.go).


## Contributions

This is an open source software licensed under Apache License 2.0.

Feel free to open issues and pull requests for improving code and documents.
