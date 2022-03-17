package main

import (
	"testing"
)

func TestPrinter(t *testing.T) {

	success := printer("Just a test")

	if success != true {
		t.Errorf("The test failed")
	}
}
