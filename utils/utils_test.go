package utils

import (
	"bytes"
	"testing"
)

var wrongformataddress = []string{
	"0x3", "", "2435",
	"27d22890587cfada7fec247c5180d73de6c670c4",
	"0xdj2890587cfada7fec247c5180d73de6c670c4",
	"027d22890587cfada7fec247c5180d73de6c670c4",
	"0027d22890587cfada7fec247c5180d73de6c670c4",
	"xx27d22890587cfada7fec247c5180d73de6c670c4",
	"XX27d22890587cfada7fec247c5180d73de6c670c4",
}

var correctformataddress = []string{
	"0x27d22890587cfada7fec247c5180d73de6c670c4",
	"0X27d22890587cfada7fec247c5180d73de6c670c4",
	"0x27D22890587CFADA7FEC247C5180D73DE6C670C4",
	"0X27D22890587CFADA7FEC247C5180D73DE6C670C4",
}

var addressbytes = []byte{39, 210, 40, 144, 88, 124, 250, 218,
	127, 236, 36, 124, 81, 128, 215, 61, 230, 198, 112, 196}

func TestIsValidAddress(t *testing.T) {
	for _, wrongvalue := range wrongformataddress {
		isvalid := IsValidAddress(wrongvalue)
		expected := false
		if isvalid != expected {
			t.Errorf("got %v, wanted %v", isvalid, expected)
		}
	}

	for _, correctvalue := range correctformataddress {
		isvalid := IsValidAddress(correctvalue)
		expected := true
		if isvalid != expected {
			t.Errorf("got %v, wanted %v", isvalid, expected)
		}
	}
}

func TestAddresstoBytes(t *testing.T) {
	for _, wrongvalue := range wrongformataddress {
		var isvalid, expected bool
		_, err := AddresstoBytes(wrongvalue)
		if err != nil {
			isvalid = false
		}
		expected = false
		if isvalid != expected {
			t.Errorf("got %v, wanted %v", isvalid, expected)
		}
	}

	for _, correctvalue := range correctformataddress {
		var isvalid, expected bool
		convaddressbytes, err := AddresstoBytes(correctvalue)
		if err == nil && bytes.Compare(convaddressbytes, addressbytes) == 0 {
			isvalid = true
		}
		expected = true
		if isvalid != expected {
			t.Errorf("got %v, wanted %v", isvalid, expected)
		}
	}
}
