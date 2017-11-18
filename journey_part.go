package ns

import "encoding/xml"

// JourneyPart is a single ride
type JourneyPart struct {
	XMLName      xml.Name      `xml:"ReisDeel"`
	Agency       string        `xml:"Vervoerder"`
	VehicleType  string        `xml:"VervoerType"`
	VerhicleSpec int           `xml:"RitNummer"`
	Status       string        `xml:"Status"`
	JourneyStops []JourneyStop `xml:"ReisStop"`
}
