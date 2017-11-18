package ns

import (
	"context"
	"net/http"
	"reflect"
	"testing"
	"time"
)

const (
	FromStation = "Amsterdam Centraal"
	ToStation   = "Den Haag Centraal"
	NextAdvices = 5
)

func TestBasicTravelOptionsRequest(t *testing.T) {
	ns := Init(Username, Password)

	ctx := context.TODO()
	req := &TravelOptionsRequest{
		From:        FromStation,
		To:          ToStation,
		NextAdvices: NextAdvices,
		DateTime:    time.Now(),
	}
	options, res, err := ns.TravelOptions.Get(ctx, req)

	expect(t, err, nil)
	expect(t, res.Response.StatusCode, http.StatusOK)
	expect(t, reflect.TypeOf(options), reflect.TypeOf(TravelOptions{}))
	expect(t, res.Success, true)
}
