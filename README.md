# go-yahoo-weather [![GoDoc](https://godoc.org/github.com/int128/go-yahoo-weather/weather?status.svg)](https://godoc.org/github.com/int128/go-yahoo-weather/weather) [![CircleCI](https://circleci.com/gh/int128/go-yahoo-weather.svg?style=shield)](https://circleci.com/gh/int128/go-yahoo-weather)

This is a [Yahoo Weather API](https://developer.yahoo.co.jp/webapi/map/openlocalplatform/v1/weather.html) Client for Go.

See [GoDoc](https://godoc.org/github.com/int128/go-yahoo-weather/weather).


## TL;DR

You need to get a Client ID on [Yahoo Developers Network](https://developer.yahoo.co.jp).

For example,

```go
package main

import (
	"github.com/int128/go-yahoo-weather/weather"
	"log"
	"os"
)

func main() {
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
```


## Contributions

This is an open source software licensed under Apache License 2.0.

Feel free to open issues and pull requests for improving code and documents.
