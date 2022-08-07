package erc20

import (
	"golang.org/x/crypto/sha3"
)

// ERC20-based token standard; getters and functions.
var (
	name         string = "name()"
	symbol       string = "symbol()"
	decimals     string = "decimals()"
	totalSupply  string = "totalSupply()"
	balance      string = "balanceOf(address)"
	transfer     string = "transfer(address,uint256)"
	approve      string = "approve(address,uint256)"
	transferFrom string = "transferFrom(address,address,uint256)"
	allowance    string = "allowance(address,address)"
)

// ER20TokenDecimals holds the decimal placing value of the ERC20-based token standard.
var ER20TokenDecimals = make(map[string]uint8)

// erc20methodIDs holds the method ID of the ERC20-based token standard.
var erc20methodIDs = make(map[string][4]byte)

func init() {
	erc20standard := []string{name, symbol, decimals, totalSupply, balance,
		transfer, approve, transferFrom, allowance}

	for _, method := range erc20standard {
		erc20methodIDs[method] = methodID(method)
	}
}

// methodID calculates and returns the method ID of 4 bytes.
func methodID(method string) [4]byte {
	methodbytes := []byte(method)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(methodbytes)

	var methodSignature [4]byte
	for i, id := range hash.Sum(nil)[:4] {
		methodSignature[i] = id
	}
	return methodSignature
}
