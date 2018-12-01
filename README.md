# go-yahoo-weather [![GoDoc](https://godoc.org/github.com/int128/go-yahoo-weather/weather?status.svg)](https://godoc.org/github.com/int128/go-yahoo-weather/weather) [![CircleCI](https://circleci.com/gh/int128/go-yahoo-weather.svg?style=shield)](https://circleci.com/gh/int128/go-yahoo-weather)

This is a Go package for [Yahoo Japan Weather API](https://developer.yahoo.co.jp/webapi/map/openlocalplatform/v1/weather.html).

See [GoDoc](https://godoc.org/github.com/int128/go-yahoo-weather/weather).


## Getting Started

You need to get a Client ID on [Yahoo Japan Developers Network](https://developer.yahoo.co.jp).

You can get weather observation and forecast as follows:

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

Response struct should be as follows:

```
{Body:{
  ResultInfo:{Count:1 Total:1 Start:1 Status:200 Latency:0.005317 Description: Copyright:(C) Yahoo Japan Corporation.}
  Feature:[
   {ID:201812011520_139.73229_35.663613
    Name:地点(139.73229,35.663613)の2018年12月01日 15時20分から60分間の天気情報
    Geometry:{Type:point Coordinates:139.73229,35.663613}
    Property:{
     WeatherAreaCode:4410
     WeatherList:{Weather:[
      {Type:observation Date:201812011520 Rainfall:0}
      {Type:forecast Date:201812011530 Rainfall:0}
      {Type:forecast Date:201812011540 Rainfall:0}
      {Type:forecast Date:201812011550 Rainfall:0}
      {Type:forecast Date:201812011600 Rainfall:0}
      {Type:forecast Date:201812011610 Rainfall:0}
      {Type:forecast Date:201812011620 Rainfall:0}
     ]}}}]}
 Expires:2018-12-01 06:34:38 +0000 UTC}
```


## Contributions

This is an open source software licensed under Apache License 2.0.

Feel free to open issues and pull requests for improving code and documents.
