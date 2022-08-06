package erc20

import (
	"reflect"
	"testing"
)

var (
	address         string = "0x27d22890587cfada7fec247c5180d73de6c670c4"
	contractaddress string = "0xf6fe970533fe5C63d196139B14522Eb2956f8621"
)

func createERC20Clause() (*ERC20Clause, error) {
	erc20clause, err := New().
		AddToAddress(address).
		AddValue("3").
		AddTokenAddress(contractaddress).
		Build()
	return erc20clause, err
}

func TestCreateClause(t *testing.T) {
	erc20clause, err := createERC20Clause()
	if err != nil {
		t.Errorf("cannot create erc20clause: %v", err)
	}

	expectederc20clause := &ERC20Clause{
		ERC20Body{
			to:           address,
			value:        "3",
			tokenAddress: contractaddress,
		},
	}

	if !reflect.DeepEqual(erc20clause, expectederc20clause) {
		t.Errorf("got %v, wanted %v", erc20clause, expectederc20clause)
	}
}
