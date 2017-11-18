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
