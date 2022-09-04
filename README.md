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
