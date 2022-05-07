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

var correctdecimalvalue = []string{
	"3", "78", "99", "215", "1000000000000000000",
}

var wrongdecimalvalue = []string{
	"0x3", "", "3.3", "0.2", "0xff", "00.5", "55.55",
	"00.001", "efb2", "8.007", "70.007", "666.666",
}

func TestIsValidDecimalValue(t *testing.T) {
	for _, wrongvalue := range wrongdecimalvalue {
		var isvalid, expected bool
		isvalid = IsValidDecimalValue(wrongvalue)
		expected = false
		if isvalid != expected {
			t.Errorf("got %v, wanted %v", isvalid, expected)
		}
	}

	for _, correctvalue := range correctdecimalvalue {
		var isvalid, expected bool
		isvalid = IsValidDecimalValue(correctvalue)
		expected = true
		if isvalid != expected {
			t.Errorf("got %v, wanted %v", isvalid, expected)
		}
	}
}

var wrongvalues = []string{
	"0x3", "", "00.00.1", "10.2.2", "0xff", "efb2",
}

var correctvalues = []string{
	"3", "0.999", "5.2", "2", "1.002", "00.5", "88.95",
	"00.001", "9.9", "8.007", "0.007", "9045.5", "55.55",
}

func TestIsValidValue(t *testing.T) {
	for _, wrongvalue := range wrongvalues {
		var isvalid, expected bool
		isvalid = IsValidValue(wrongvalue)
		expected = false
		if isvalid != expected {
			t.Errorf("got %v, wanted %v", isvalid, expected)
		}
	}

	for _, correctvalue := range correctvalues {
		var isvalid, expected bool
		isvalid = IsValidValue(correctvalue)
		expected = true
		if isvalid != expected {
			t.Errorf("got %v, wanted %v", isvalid, expected)
		}
	}
}

func TestLeftPadBytes(t *testing.T) {
	var isvalid, expected bool

	var leftpadedbytes = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 39, 210, 40, 144, 88, 124, 250, 218, 127, 236,
		36, 124, 81, 128, 215, 61, 230, 198, 112, 196}

	resultbytes := LeftPadBytes(addressbytes, 32)
	if bytes.Compare(resultbytes, leftpadedbytes) == 0 {
		isvalid = true
	}
	expected = true
	if isvalid != expected {
		t.Errorf("got %v, wanted %v", isvalid, expected)
	}
}
