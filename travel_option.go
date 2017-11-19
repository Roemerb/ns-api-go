package ns

import (
	"context"
	"encoding/xml"
	"time"
)

// TravelOptionEndpoint is the endpoint for the travel options API
const TravelOptionEndpoint = "ns-api-treinplanner"

// TravelOptions is the response format from the NS API
type TravelOptions struct {
	XMLName xml.Name       `xml:"NS"`
	Options []TravelOption `xml:"ReisMogelijkheden>ReisMogelijkheid"`
}

// TravelOption is a series of connections to the destination
type TravelOption struct {
	XMLName                xml.Name      `xml:"ReisMogelijkheid"`
	NumberOfTransfers      int           `xml:"AantalOverstappen"`
	ExpectedTravelTime     string        `xml:"GeplandeReisTijd"`
	ActualTravelTime       string        `xml:"ActueleReisTijd"`
	Optimal                bool          `xml:"Optimaal"`
	ScheduledDepartureTime string        `xml:"GeplandeVertrekTijd"`
	CurrentDepartureTime   string        `xml:"ActueleVertrekTijd"`
	ScheduledArrivalTime   string        `xml:"GeplandeAankomstTijd"`
	CurrentArrivalTime     string        `xml:"ActueleAankomstTijd"`
	Status                 string        `xml:"Status"`
	JourneyType            string        `xml:"reisSoort,attr"`
	JourneyParts           []JourneyPart `xml:"ReisDeel"`
}

// TravelOptionsRequest is the format of a request for travel options
type TravelOptionsRequest struct {
	From              string    `url:"fromStation"`
	To                string    `url:"toStation"`
	Via               string    `url:"viaStation"`
	NextAdvices       int       `url:"nextAdvices"`
	PreviousAdvices   int       `url:"previousAdvices"`
	DateTime          time.Time `url:"dateTime"`
	DateTimeIsArrival bool      `url:"Departure"`
	HSLAllowed        bool      `url:"hslAllowed"`
	HasYearCard       bool      `url:"yearCard"`
}

// TravelOptionService describes the methods for the TravelOptionsService
type TravelOptionService interface {
	Get(ctx context.Context) (TravelOptions, APIResponse, error)
}

// TravelOptionServiceImpl implements the TravelOptionsService
type TravelOptionServiceImpl struct {
	ns *NS
}

// Get will execute a TravelOptionsRequest
func (tos *TravelOptionServiceImpl) Get(ctx context.Context, req *TravelOptionsRequest) (TravelOptions, APIResponse, error) {
	var apiResponse APIResponse
	var options TravelOptions

	res, err := tos.ns.DoRequest(ctx, TravelOptionEndpoint, req, true)
	apiResponse.Response = res
	if err != nil {
		return options, apiResponse, err
	}
	err = tos.ns.ParseResponse(res, &options, true)
	if err != nil {
		return options, apiResponse, err
	}

	apiResponse.Result = &options
	apiResponse.Success = true
	return options, apiResponse, nil
}

// GetDelayInMinutes calculates the total delay in minutes of all the stops
func (option TravelOption) GetDelayInMinutes() int {
	delay := 0

	for _, part := range option.JourneyParts {
		for _, stop := range part.JourneyStops {
			delay = delay + stop.GetDelayInMinutes()
		}
	}

	return delay
}
