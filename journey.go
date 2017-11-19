package ns

import (
	"context"
	"encoding/xml"
)

// CurrentJourneysEndpoint is the endpoint for the departures API
const CurrentJourneysEndpoint = "ns-api-avt"

// Journeys is a collection of journeys
type Journeys struct {
	XMLName  xml.Name  `xml:"NS"`
	Journeys []Journey `xml:"ActueleVertrekTijden>VertrekkendeTrein"`
}

// Journey is a departing train from a station from the departures API
type Journey struct {
	JourneyCode           int    `xml:"RitNummer"`
	DepartureTime         string `xml:"VertrekTijd"`
	DepartureDelay        string `xml:"VertrekVertraging"`
	DepartureDelayText    string `xml:"VertrekVertragingTekst"`
	RouteText             string `xml:"RouteTekst"`
	Destination           string `xml:"EindBestemming"`
	TrainType             string `xml:"TreinSoort"`
	TransportationCompany string `xml:"Vervoerder"`
	Track                 string `xml:"VertrekSpoor"`
}

// JourneyService describes the methods in the Journeys service
type JourneyService interface {
	GetForStation(ctx context.Context, station Station) (Journeys, APIResponse, error)
	GetForStationCode(ctx context.Context, code int) (Journeys, APIResponse, error)
}

// JourneyServiceImpl implements the JourneyService and provides an NS instance
type JourneyServiceImpl struct {
	ns *NS
}

// JourneysRequest is the format for a request to the departures API
type JourneysRequest struct {
	StationCode string `url:"station"`
}

// GetForStation gets the current departures for a station
func (jsi *JourneyServiceImpl) GetForStation(ctx context.Context, station Station) (Journeys, APIResponse, error) {
	return jsi.get(ctx, JourneysRequest{StationCode: station.Code})
}

// GetForStationCode gets the current departures for a station by it's code
func (jsi *JourneyServiceImpl) GetForStationCode(ctx context.Context, code string) (Journeys, APIResponse, error) {
	return jsi.get(ctx, JourneysRequest{StationCode: code})
}

func (jsi JourneyServiceImpl) get(ctx context.Context, request interface{}) (Journeys, APIResponse, error) {
	var apiResponse APIResponse
	var journeys Journeys

	res, err := jsi.ns.DoRequest(ctx, CurrentJourneysEndpoint, request, true)
	apiResponse.Response = res
	if err != nil {
		return Journeys{}, apiResponse, err
	}
	err = jsi.ns.ParseResponse(res, &journeys, true)
	if err != nil {
		return Journeys{}, apiResponse, err
	}

	apiResponse.Result = &journeys
	apiResponse.Success = true
	return journeys, apiResponse, nil
}
