package controllers

import (
	"testing"
)

func TestDerivePaymentAddress(t *testing.T) {
	salt := "payment-test"
	addr, _, err := derivePaymentAddress("0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6", "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC", salt)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if addr.Hex() != "0x756AC28d5Dd66E0fc3b79A7aC1d3184436524E6E" {
		t.Errorf("Address mismatch: %v", addr.Hex())
	}
}
