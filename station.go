package ns

import (
	"context"
	"encoding/xml"
)

// StationEndpoint is the endpoint of the station v2 API
const StationEndpoint = "ns-api-stations-v2"

// Stations the response format of the stations v2 API
type Stations struct {
	XMLName  xml.Name  `xml:"NS"`
	Stations []Station `xml:"Stations>Station"`
}

// Station is the format of a station
type Station struct {
	Code  string `xml:"Code"`
	Type  string `xml:"Type"`
	Names struct {
		Short  string `xml:"Kort"`
		Middle string `xml:"Middel"`
		Long   string `xml:"Lang"`
	}
	Country  string           `xml:"Land"`
	UICCode  int              `xml:"UICCode"`
	Lat      float64          `xml:"Lat"`
	Lon      float64          `xml:"Lon"`
	Synonyms []StationSynonym `xml:"Synoniemen"`
}

// StationSynonym is a synonymical name for a station
type StationSynonym struct {
	Synonym string `xml:"Synonym"`
}

// StationService describes the methods in the Station Service
type StationService interface {
	Get(ctx context.Context) (Stations, APIResponse, error)
}

// StationServiceImpl implements the Station Service and provides an NS instance
type StationServiceImpl struct {
	ns *NS
}

// Get uses the stations v2 API to fetch all available stations
func (service StationServiceImpl) Get(ctx context.Context) (Stations, APIResponse, error) {
	var apiResponse APIResponse
	var stations Stations

	res, err := service.ns.DoRequest(ctx, StationEndpoint, nil, true)
	apiResponse.Response = res
	if err != nil {
		return stations, apiResponse, err
	}
	err = service.ns.ParseResponse(res, &stations, true)
	if err != nil {
		return stations, apiResponse, err
	}

	apiResponse.Result = &stations
	apiResponse.Success = true
	return stations, apiResponse, nil
}
