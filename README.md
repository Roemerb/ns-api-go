# NS Go API Library [![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/roemerb/ns-api-go)

## Summary

This is an API library for the Dutch Railway Service (NS) written in Go. Use of the library requires API credentials. These can be requested [here](https://www.ns.nl/ews-aanvraagformulier/?0).

## Installation

```sh
go get github.com/roemerb/ns-api-go
```

## Documentation

The use of this library starts with initialisation:

```go
import "github.com/roemerb/ns-api-go"

const (
	Username = "you@example.com"
	Password = "YourAPIKey"
)

ns := ns.Init(Username, Password)
```

### Travel options

The Travel Options API provides detailed guidance from and to a station. It consists of multiple `JourneyPart`'s, each representing a single train ride.

To get travel options, create a `TravelOptionsRequest` describing your journey:

```go
req := TravelOptionsRequest{
	From: "Amsterdam Centraal",
	To: "Rotterdam Centraal",
	Via: "Schiphol"
	// Number of advices in the future. To get advices in the past, use PreviousAdvices
	NextAdvices: 5,
	DateTime: time.Now(),
	// Set this to true to indicate that DateTime is the time you want to arrive. False for departure time
	DateTimeIsArrival: true,
	HSLAllowed: true, // Allow high speed trains (higher fare)
}
```

Now use the request to fetch the travel options:

```go
options, apiResponse, err := ns.TravelOptions.Get(context.TODO(), req)
```

### Stations

The NS also provides a stations API. This API will contain details about all train stations in the Netherlands. It'll also contain the station codes. These codes can be used to fetch all departing trains from a station.

To get a list of stations:

```go
stations, apiResponse, err := ns.Stations.Get(context.TODO())

// Print the names and codes of stations
for _, station := range stations.Stations {
	fmt.Println(station.Names.Middle + ": " + station.Code)
}
```

For more details on the use of the other API's, refer to the test files or the [GoDoc](http://godoc.org/github.com/roemerb/ns-api-go).