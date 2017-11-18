package ns

import "encoding/xml"

// JourneyStop is a stop in a ride
type JourneyStop struct {
	XMLName        xml.Name `xml:"ReisStop"`
	StationName    string   `xml:"Naam"`
	DepartureDelay string   `xml:"VertrekVertraging"`
	ArrivalDelay   string   `xml:"AankomstVertraging"`
	Time           string   `xml:"Tijd"`
	Track          string   `xml:"Spoor"`
	TrackChanged   bool     `xml:"wijziging,attr"`
}
