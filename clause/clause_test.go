package clause

import (
	"reflect"
	"testing"
)

var (
	address string = "0x27d22890587cfada7fec247c5180d73de6c670c4"
)

func TestCreateClause(t *testing.T) {
	clause, err := New().AddToAddress(address).AddValue("2").Build()
	if err != nil {
		t.Errorf("cannot create clause: %v", err)
	}

	expectedclause := &Clause{
		ClauseBody{
			to:    address,
			value: "2",
		},
	}

	if !reflect.DeepEqual(clause, expectedclause) {
		t.Errorf("got %v, wanted %v", clause, expectedclause)
	}
}
