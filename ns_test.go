package ns

import (
	"reflect"
	"testing"
)

const (
	Username = ""
	Password = ""
)

func TestInitiatesProperly(t *testing.T) {
	ns := Init(Username, Password)

	expect(t, reflect.TypeOf(ns), reflect.TypeOf(&NS{}))
}

func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected [%v] (type %v) - Got [%v] (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}
