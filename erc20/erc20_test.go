package erc20

import (
	"encoding/hex"
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

func TestName(t *testing.T) {
	erc20clause, err := createERC20Clause()
	if err != nil {
		t.Errorf("cannot create erc20clause: %v", err)
	}

	payloaddata := erc20clause.TokenName()
	hexvaluepayload := hex.EncodeToString(payloaddata)
	expected := "06fdde03"

	if hexvaluepayload != expected {
		t.Errorf("got %v, wanted %v", payloaddata, expected)
	}
}

func TestSymbol(t *testing.T) {
	erc20clause, err := createERC20Clause()
	if err != nil {
		t.Errorf("cannot create erc20clause: %v", err)
	}

	payloaddata := erc20clause.TokenSymbol()
	hexvaluepayload := hex.EncodeToString(payloaddata)
	expected := "95d89b41"

	if hexvaluepayload != expected {
		t.Errorf("got %v, wanted %v", hexvaluepayload, expected)
	}
}

func TestDecimals(t *testing.T) {
	erc20clause, err := createERC20Clause()
	if err != nil {
		t.Errorf("cannot create erc20clause: %v", err)
	}

	payloaddata := erc20clause.TokenDecimals()
	hexvaluepayload := hex.EncodeToString(payloaddata)
	expected := "313ce567"

	if hexvaluepayload != expected {
		t.Errorf("got %v, wanted %v", hexvaluepayload, expected)
	}
}

func TestTotalSupply(t *testing.T) {
	erc20clause, err := createERC20Clause()
	if err != nil {
		t.Errorf("cannot create erc20clause: %v", err)
	}

	payloaddata := erc20clause.TokenTotalSupply()
	hexvaluepayload := hex.EncodeToString(payloaddata)
	expected := "18160ddd"

	if hexvaluepayload != expected {
		t.Errorf("got %v, wanted %v", hexvaluepayload, expected)
	}
}

func TestBalance(t *testing.T) {
	erc20clause, err := createERC20Clause()
	if err != nil {
		t.Errorf("cannot create erc20clause: %v", err)
	}

	payloaddata, err := erc20clause.TokenBalance()
	if err != nil {
		t.Errorf("cannot create payload data for token balance: %v", err)
	}

	hexvaluepayload := hex.EncodeToString(payloaddata)
	expected := "70a0823100000000000000000000000027d22890587cfada7fec247c5180d73de6c670c4"
	if hexvaluepayload != expected {
		t.Errorf("got %v, wanted %v", hexvaluepayload, expected)
	}
}

func TestTransfer(t *testing.T) {
	erc20clause, err := createERC20Clause()
	if err != nil {
		t.Errorf("cannot create erc20clause: %v", err)
	}

	erc20clause.AddValue("50000000000000000000")
	payloaddata, err := erc20clause.TokenTranfer()
	if err != nil {
		t.Errorf("cannot create payload data for token transfer: %v", err)
	}

	hexvaluepayload := hex.EncodeToString(payloaddata)
	expected := "a9059cbb00000000000000000000000027d22890587cfada7fec247c5180d73de6c670c4000000000000000000000000000000000000000000000002b5e3af16b1880000"
	if hexvaluepayload != expected {
		t.Errorf("got %v, wanted %v", hexvaluepayload, expected)
	}
}
