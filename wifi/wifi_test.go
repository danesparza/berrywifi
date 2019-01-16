package wifi_test

import (
	"testing"

	"github.com/danesparza/berrywifi/wifi"
)

/**
Example output:
*/
var output = []byte(` 
	SSID: Chupacabra
	SSID: Princess_Wifi
	SSID: Chupacabra
	SSID: Nahual
	SSID:
	SSID:
	SSID: Nahual
	SSID: Princess_Wifi
	SSID: ATT7M2Y587
	SSID: MarksPlace
	SSID: ATT34AH43d
	SSID: HOME-F862
	SSID: Default-guest
	SSID: ATT0500
	SSID: Zapeda
	Extended capabilities: TFS, WNM-Sleep Mode, TIM Broadcast, BSS Transition, SSID List, 6
	SSID: xfinitywifi
`)

func TestGetAPList_ValidRequest_Successful(t *testing.T) {
	//	Arrange

	//	Act
	aps, err := wifi.GetAPList(string(output))

	//	Assert
	if err != nil {
		t.Errorf("GetAPList - should handle request successfully, but got error: %v", err)
	}

	if len(aps) < 1 {
		t.Errorf("GetAPList - should return a list of APs but got none")
	}

	t.Logf("The list of aps: %v", aps)
}
