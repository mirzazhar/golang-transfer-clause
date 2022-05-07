package utils

import "errors"

var ErrTokenAddress = errors.New("token address format is invalid or nil")
var ErrToAddress = errors.New("recipient account address format is invalid or nil")
var ErrDecimalPoint = errors.New("decimal point must not be a negative number")
var ErrValue = errors.New("value must be non-empty e.g. integer or decimal point number")
var ErrDecimalValue = errors.New("the value must be given as an integer without a decimal point")
var ErrSameEOAContractAddr = errors.New("externally onwed address (EOA) and contract address can never b same")
var ErrAddressLength = errors.New("invalid address length; it must be 40 (without prefix 0x) or 42 (with prefix 0x)")
