package ns

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestGetStations(t *testing.T) {
	ns := Init(Username, Password)

	ctx := context.TODO()
	stations, res, err := ns.Stations.Get(ctx)

	expect(t, err, nil)
	expect(t, res.Success, true)
	expect(t, res.Response.StatusCode, http.StatusOK)
	expect(t, reflect.TypeOf(stations), reflect.TypeOf(Stations{}))
	expect(t, len(stations.Stations) > 0, true)
}

func TestGetStationByCode(t *testing.T) {
	ns := Init(Username, Password)

	ctx := context.TODO()
	stations, _, _ := ns.Stations.Get(ctx)

	asd := stations.GetStationByCode("asd")

	expect(t, asd.Names.Long, "Amsterdam Centraal")
	expect(t, asd.Code, "ASD")
}
