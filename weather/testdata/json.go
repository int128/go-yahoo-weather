package testdata

const WeatherResponseJSON = `
{
  "ResultInfo": {
    "Count": 2,
    "Total": 2,
    "Start": 1,
    "Status": 200,
    "Latency": 0.005322,
    "Description": "",
    "Copyright": "(C) Yahoo Japan Corporation."
  },
  "Feature": [
    {
      "Id": "201812121305_139.73229_35.663613",
      "Name": "地点(139.73229,35.663613)の2018年12月12日 13時05分から60分間の天気情報",
      "Geometry": {
        "Type": "point",
        "Coordinates": "139.73229,35.663613"
      },
      "Property": {
        "WeatherAreaCode": 4410,
        "WeatherList": {
          "Weather": [
            {
              "Type": "observation",
              "Date": "201812121305",
              "Rainfall": 0.00
            },
            {
              "Type": "forecast",
              "Date": "201812121315",
              "Rainfall": 0.00
            },
            {
              "Type": "forecast",
              "Date": "201812121325",
              "Rainfall": 0.00
            },
            {
              "Type": "forecast",
              "Date": "201812121335",
              "Rainfall": 0.00
            },
            {
              "Type": "forecast",
              "Date": "201812121345",
              "Rainfall": 0.00
            },
            {
              "Type": "forecast",
              "Date": "201812121355",
              "Rainfall": 0.00
            },
            {
              "Type": "forecast",
              "Date": "201812121405",
              "Rainfall": 0.00
            }
          ]
        }
      }
    },
    {
      "Id": "201812121305_140.72892_41.768674",
      "Name": "地点(140.72892,41.768674)の2018年12月12日 13時05分から60分間の天気情報",
      "Geometry": {
        "Type": "point",
        "Coordinates": "140.72892,41.768674"
      },
      "Property": {
        "WeatherAreaCode": 2300,
        "WeatherList": {
          "Weather": [
            {
              "Type": "observation",
              "Date": "201812121305",
              "Rainfall": 0.35
            },
            {
              "Type": "forecast",
              "Date": "201812121315",
              "Rainfall": 0.45
            },
            {
              "Type": "forecast",
              "Date": "201812121325",
              "Rainfall": 1.15
            },
            {
              "Type": "forecast",
              "Date": "201812121335",
              "Rainfall": 0.45
            },
            {
              "Type": "forecast",
              "Date": "201812121345",
              "Rainfall": 1.85
            },
            {
              "Type": "forecast",
              "Date": "201812121355",
              "Rainfall": 0.00
            },
            {
              "Type": "forecast",
              "Date": "201812121405",
              "Rainfall": 0.00
            }
          ]
        }
      }
    }
  ]
}
`
