package lockerdome

import (
	"encoding/json"
	"testing"

	"github.com/prebid/prebid-server/openrtb_ext"

	"fmt"
)

// This file actually intends to test static/bidder-params/lockerdome.json
//
// These also validate the format of the external API: request.imp[i].ext.lockerdome

// TestValidParams makes sure that the LockerDome schema accepts all imp.ext fields which we intend to support.
func TestValidParams(t *testing.T) {
	fmt.Println("---------testing valid params")
	validator, err := openrtb_ext.NewBidderParamsValidator("../../static/bidder-params")
	if err != nil {
		t.Fatalf("Failed to fetch the json-schemas. %v", err)
	}

	for _, validParam := range validParams {
		if err := validator.Validate(openrtb_ext.BidderLockerDome, json.RawMessage(validParam)); err != nil {
			t.Errorf("Schema rejected LockerDome params: %s", validParam)
		}
	}
}


// TestInvalidParams makes sure that the LockerDome schema rejects all the imp.ext fields we don't support.
func TestInvalidParams(t *testing.T) {
	fmt.Println("---------testing invalid params")
	validator, err := openrtb_ext.NewBidderParamsValidator("../../static/bidder-params")
	if err != nil {
		t.Fatalf("Failed to fetch the json-schemas. %v", err)
	}

	for _, invalidParam := range invalidParams {
		if err := validator.Validate(openrtb_ext.BidderLockerDome, json.RawMessage(invalidParam)); err == nil {
			t.Errorf("Schema allowed unexpected params: %s", invalidParam)
		}
	}
}

// TODO: string vs number?
var validParams = []string{
	`{"adUnitId": "1234567890"}` // adUnitID can be a string of numbers
	`{"adUnitId": "LD1234567890"}`, // adUnitId can start with "LD"
}

var invalidParams = []string{
	``,
	`null`,
	`true`,
	`1`,
	`1.5`,
	`[]`,
	`{}`,
	`{"adUnitId": "LD"}`, // adUnitId can't just be "LD"
	`{"adUnitId": 1234567890}` // adUnitID can't be a number
}
