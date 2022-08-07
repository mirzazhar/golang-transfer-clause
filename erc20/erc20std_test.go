package erc20

import (
	"encoding/hex"
	"testing"
)

func TestMethodID(t *testing.T) {
	methodIDs := make(map[string]string)
	methodIDs[name] = "06fdde03"
	methodIDs[symbol] = "95d89b41"
	methodIDs[decimals] = "313ce567"
	methodIDs[totalSupply] = "18160ddd"
	methodIDs[balance] = "70a08231"
	methodIDs[transfer] = "a9059cbb"
	methodIDs[approve] = "095ea7b3"
	methodIDs[transferFrom] = "23b872dd"
	methodIDs[allowance] = "dd62ed3e"

	erc20standard := []string{name, symbol, decimals, totalSupply, balance,
		transfer, totalSupply, approve, transferFrom, allowance}

	for _, method := range erc20standard {
		id := methodID(method)
		methodId := hex.EncodeToString(id[:])
		expectedmethodID := methodIDs[method]

		if methodId != expectedmethodID {
			t.Errorf("got %v, wanted %v", methodId, expectedmethodID)
		}
	}
}
