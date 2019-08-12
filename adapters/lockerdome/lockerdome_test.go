package lockerdome

import (
	"github.com/prebid/prebid-server/adapters/adapterstest"
	"testing"
)

func TestJsonSamples(t *testing.T) {
	adapterstest.RunJSONBidderTest(t, "lockerdometest", NewLockerDomeBidder("https://local.lockerdome.com:3000/ladbid/prebidserver/openrtb2"))
}
