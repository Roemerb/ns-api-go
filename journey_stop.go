package ns

import (
	"encoding/xml"
	"strconv"
	"strings"
)

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

// GetDelayInMinutes checks if there is delay based on DepartureDelay, and converts it to an int
func (stop JourneyStop) GetDelayInMinutes() int {
	delay := 0
	if stop.DepartureDelay != "" {
		// Format is "+n min", so we'll strip the '+' and split by space
		parts := strings.Split(stop.DepartureDelay[1:], " ")
		// If the first part can be converted to an int, we got the amounts of minutes.
		// We'll add it to the total
		if minutes, err := strconv.Atoi(parts[0]); err == nil {
			delay = delay + minutes
		}
	}

	return delay
}
