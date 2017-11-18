package ns

import (
	"context"
	"encoding/xml"
	"io/ioutil"
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

	res, err := jsi.ns.DoRequest(ctx, CurrentJourneysEndpoint, request, true)
	defer res.Body.Close()
	if err != nil {
		apiResponse.Success = false
		return Journeys{}, apiResponse, err
	}

	buff, ioerr := ioutil.ReadAll(res.Body)
	if ioerr != nil {
		apiResponse.Success = false
		return Journeys{}, apiResponse, ioerr
	}
	responseString := string(buff)
	responseString = "<NS>" + responseString + "</NS>"
	var target Journeys
	var apiErr APIError
	err = xml.Unmarshal([]byte(responseString), &target)
	if err != nil {
		xml.Unmarshal([]byte(responseString), &apiErr)
		apiResponse.Result = &apiErr
		apiResponse.Success = false
		return Journeys{}, apiResponse, err
	}
	apiResponse.Success = true
	apiResponse.Response = res
	apiResponse.Result = &target

	return target, apiResponse, nil
}
