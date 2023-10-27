# Golang Transfer Clause

This package provides a toolbox for setting up the transfer object, later to be used in the transaction creation of an ethereum or fork-based transaction. 
Here, the transfer object refers to the following information:
- Receiver Address
- The amount to be transferred
- Any arbitrary data

Typically, any wallet backend, including ethereum-based forks, dapps, etc., is intended to use this package. The motivation for developing this package is to make it easier to interact with tokens such as ERC-20. 

## Features
This package provides the following features:
- Prepares the Transfer Clause Information required to transfer native ethereum or fork-based coins.
- Prepares Payload Data to be used in the Transfer Clause to interact with deployed ERC-20-based tokens in an ethereum or fork-based network setting:
  - ERC20 Token Name
  - ERC20 Token Symbol
  - ERC20 Token Decimals
  - ERC20 Token TotalSupply
  - ERC20 Token BalanceOf
  - ERC20 Token Transfer
  - ERC20 Token Approve
  - ERC20 Token Transferfrom
  - ERC20 Token Allowance
- Provides handy utility functions:
  - It converts ethereum addresses into bytes.
  - It validates the ethereum address formats. 
  - It validates the amount to be transferred in string format:
    - Verifies that the amount is provided as an integer value only.
    - Verifies that the amount is provided as an integer value or with any possible decimal point.
  - It places the specified number of zeros on the left side of the byte array.

## Disclaimer
Other popular approaches for interacting with ERC-based token smart contracts exist as well. This package provides an alternative mechanism for interacting with ERC-based tokens statically. 

**Notably, for the above-described ERC-20-based token standard, the payload preparation functionality provided by this package works well. In case of interacting with any modified/enhanced version of the ERC-20-based token, the payload preparation functionality also needs to be modified to make it compatible.**

## Installation
```sh
go get -u github.com/mirzazhar/golang-transfer-clause
```
## Usage

### Transfer Clause
```go
package main

import (
	"fmt"

	"github.com/mirzazhar/golang-transfer-clause/clause"
)

func main() {
	var address string = "0x27d22890587cfada7fec247c5180d73de6c670c4"

	clause, err := clause.
		New().
		AddToAddress(address).
		AddValue("0.5").
		Build()
	if err != nil {
		fmt.Printf("cannot create clause: %v", err)
	}
	fmt.Println("Transfer Clause: ", clause)
}

```
### ERC-20 Based Transfer Clause
```go
package main

import (
	"fmt"

	"github.com/mirzazhar/golang-transfer-clause/erc20"
)

func main() {
	var address string = "0x27d22890587cfada7fec247c5180d73de6c670c4"
	var contractAddress string = "0xf6fe970533fe5C63d196139B14522Eb2956f8621"

	erc20Clause, err := erc20.
		New().
		AddToAddress(address).
		AddValue("1000000000000000000"). // value is given in "Wei" that is equal to 1 eth (ethereum).
		AddTokenAddress(contractAddress).
		Build()
	if err != nil {
		fmt.Printf("cannot create erc-20 based clause: %v", err)
	}
	fmt.Println("ERC-20 based Transfer Clause: ", erc20Clause)
}

```
### ERC-20 Based Payload Preparation
#### ERC-20 Token Name
```go
	namePayload, err := erc20Clause.GetERCPayloadData("name")
	if err != nil {
		fmt.Printf("cannot create pyaload for name: %v", err)
	}
	fmt.Println("erc20 name payload: ",
		hex.EncodeToString(namePayload))
```
or
```go
	namePayload := erc20Clause.TokenName()
	fmt.Println("erc20 name payload: ", hex.EncodeToString(namePayload))
```
#### ERC-20 Token Symbol
```go
	symbolPayload, err := erc20Clause.GetERCPayloadData("symbol")
	if err != nil {
		fmt.Printf("cannot create pyaload for symbol: %v", err)
	}
	fmt.Println("erc20 symbol payload: ",
		hex.EncodeToString(symbolPayload))
```
or
```go
	symbolPayload := erc20Clause.TokenSymbol()
	fmt.Println("erc20 symbol payload: ", hex.EncodeToString(symbolPayload))
```
#### ERC-20 Token TotalSupply
```go
	totalSupplyPayload, err := erc20Clause.GetERCPayloadData("totalSupply")
	if err != nil {
		fmt.Printf("cannot create pyaload for totalSupply: %v", err)
	}
	fmt.Println("erc20 totalSupply payload: ",
		hex.EncodeToString(totalSupplyPayload))
```
or
```go
	totalSupplyPayload := erc20Clause.TokenTotalSupply()
	fmt.Println("erc20 totalSupply payload: ", hex.EncodeToString(totalSupplyPayload))
```
#### ERC-20 Token Decimals
```go
	decimalsPayload, err := erc20Clause.GetERCPayloadData("decimals")
	if err != nil {
		fmt.Printf("cannot create pyaload for decimals: %v", err)
	}
	fmt.Println("erc20 decimals payload: ",
		hex.EncodeToString(decimalsPayload))
```
or
```go
	decimalsPayload := erc20Clause.TokenDecimals()
	fmt.Println("erc20 decimals payload: ", hex.EncodeToString(decimalsPayload))
```
#### ERC-20 Token BalanceOf
```go
	balancePayload, err := erc20Clause.GetERCPayloadData("balanceOf")
	if err != nil {
		fmt.Printf("cannot create pyaload for balanceOf: %v", err)
	}
	fmt.Println("erc20 balanceOf payload: ",
		hex.EncodeToString(balancePayload))
```
or
```go
	balancePayload, err := erc20Clause.TokenBalance()
	if err != nil {
		fmt.Printf("cannot create pyaload for balanceOf: %v", err)
	}
	fmt.Println("erc20 balanceOf payload: ", hex.EncodeToString(balancePayload))
```
#### ERC-20 Token Transfer
```go
	transferPayload, err := erc20Clause.GetERCPayloadData("transfer")
	if err != nil {
		fmt.Printf("cannot create pyaload for transfer: %v", err)
	}
	fmt.Println("erc20 transfer payload: ",
		hex.EncodeToString(transferPayload))
```
or
```go
	transferPayload, err := erc20Clause.TokenTranfer()
	if err != nil {
		fmt.Printf("cannot create pyaload for transfer: %v", err)
	}
	fmt.Println("erc20 transfer payload: ", hex.EncodeToString(transferPayload))
```
