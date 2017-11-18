package ns

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestGetDeparturesByStation(t *testing.T) {
	ns := Init(Username, Password)

	ctx := context.TODO()
	journeys, res, err := ns.Journeys.GetForStation(ctx, Station{Code: "ut"})

	expect(t, err, nil)
	expect(t, res.Success, true)
	expect(t, res.Response.StatusCode, http.StatusOK)
	expect(t, reflect.TypeOf(journeys), reflect.TypeOf(Journeys{}))
	expect(t, len(journeys.Journeys) > 0, true)
}

func TestGetDeparturesByStationCode(t *testing.T) {
	ns := Init(Username, Password)

	ctx := context.TODO()
	journeys, res, err := ns.Journeys.GetForStationCode(ctx, "ut")

	expect(t, err, nil)
	expect(t, res.Success, true)
	expect(t, res.Response.StatusCode, http.StatusOK)
	expect(t, reflect.TypeOf(journeys), reflect.TypeOf(Journeys{}))
	expect(t, len(journeys.Journeys) > 0, true)
}
