package utils

import (
	"encoding/hex"
	"errors"
	"regexp"
)

// AddresstoBytes converts the given Ethereum-based account address.
// It returns an error if the given address is invalid.
func AddresstoBytes(address string) ([]byte, error) {
	var data []byte
	var err error
	var addrlen int = len(address)

	if addrlen == 42 || addrlen == 40 {
		if address[0] == '0' && (address[1] == 'x' || address[1] == 'X') {
			data, err = hex.DecodeString(address[2:])
			if err != nil {
				return nil, err
			}
		} else {
			data, err = hex.DecodeString(address)
			if err != nil {
				return nil, err
			}
		}
	} else {
		return nil, errors.New("invalid address length")
	}
	return data, nil
}

// IsValidAddress validates the given Ethereum-based account address. It
// returns true if the address format is valid, and otherwise returns false.
func IsValidAddress(address string) bool {
	if address != "" && (address[:2] == "0x" || address[:2] == "0X") {
		address = address[2:]
	} else {
		return false
	}

	regexaddr := regexp.MustCompile("^[0-9a-fA-F]{40}$")
	isValidaddr := regexaddr.MatchString(address)
	return isValidaddr
}
