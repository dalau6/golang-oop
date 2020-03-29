package vin_test

import (
	"testing"
	"vin-stages/1"
)

const testVIN = "W09000051T2123456"

func TestVIN_Manufacturer(t *testing.T) {

	manufacturer := vin.Manufacturer(testVIN)
	if manufacturer != "W09123" {
		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVIN)
	}
}
