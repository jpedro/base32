package main

import (
	"testing"
)

func TestEncode(t *testing.T) {
	payload := "test"
	encoded := encode(payload)
	decoded := decode(encoded)

	if decoded != payload {
		t.Error("Expected", payload, "got", decoded)
	}
}
