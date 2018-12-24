package testdata

// curl -H "User-Agent: Yahoo AppID: CLIENT_ID" "https://map.yahooapis.jp/weather/V1/place?output=json&coordinates="
// HTTP/2 200
const ErrorResponse400JSON = `{"Error":{"Code":400,"Message":"Bad Request."}}`

// curl "https://map.yahooapis.jp/weather/V1/place?output=json&coordinates="
// HTTP/2 401
const ErrorResponse401XML = `<?xml version="1.0" encoding="utf-8" ?>
<Error>
<Message>
Bad Request: Authentication parameters in your request incompleted.
</Message>
</Error>
`

// curl "https://map.yahooapis.jp/weather/V1/place?output=json&coordinates="
// HTTP/2 403
const ErrorResponse403XML = `<?xml version="1.0" encoding="utf-8" ?>
<Error>
<Message>
Your Request was Forbidden
</Message>
</Error>
`
