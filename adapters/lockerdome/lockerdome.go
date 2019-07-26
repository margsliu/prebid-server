package lockerdome

import ( // cut down to the ones we absolutely need
	// "bytes"
	// "context"
	"encoding/json" // essential?
	"fmt" // essential?
	"net/http" // essential?
	// "io/ioutil"
	// "strconv"
	// "strings"

	// "github.com/buger/jsonparser"
	// "github.com/prebid/prebid-server/pbs"

	// "golang.org/x/net/context/ctxhttp"

	"github.com/mxmCherry/openrtb" // essential?
	"github.com/prebid/prebid-server/adapters" // essential?
	"github.com/prebid/prebid-server/errortypes" // essential?
	// "github.com/prebid/prebid-server/openrtb_ext" // might need bc of our params?
	// "github.com/prebid/prebid-server/pbsmetrics"
)

// Implements Bidder interface.
type LockerDomeAdapter struct {
	endpoint string
}

// MakeRequests makes the HTTP requests which should be made to fetch bids [from the bidder, in this case, LockerDome]
func (adapter *LockerDomeAdapter) MakeRequests(
	openRTBRequest *openrtb.BidRequest,
	extraReqInfo *adapters.ExtraRequestInfo,
) (
	requestsToBidder []*adapters.RequestData,
	errs []error,
) {


	openRTBRequestJSON, err := json.Marshal(openRTBRequest)

	if err != nil {
		errs = append(errs, err)
		return nil, errs
	}

	headers := http.Header{}
	headers.Add("Content-Type", "application/json;charset=utf-8")
	requestToBidder := &adapters.RequestData{
		Method:  "POST",
		Uri:     adapter.endpoint,
		Body:    openRTBRequestJSON,
		Headers: headers,
	}
	fmt.Println("---------TEST--------------------------------- ")
	fmt.Println("---------endpoint:\n", adapter.endpoint)
	fmt.Println("---------openRTBRequestJSON\n", string(openRTBRequestJSON))
	requestsToBidder = append(requestsToBidder, requestToBidder)


	return requestsToBidder, errs
}

const unexpectedStatusCodeMessage = "Unexpected status code: %d. Run with request.debug = 1 for more info"

// MakeBids unpacks the server's response into Bids.
// internal original request in OpenRTB, external = result of us having converted it (what comes out of MakeRequests)
func (adapter *LockerDomeAdapter) MakeBids(
	openRTBRequest *openrtb.BidRequest,
	requestToBidder *adapters.RequestData,
	bidderRawResponse *adapters.ResponseData,
) (
	bidderResponse *adapters.BidderResponse,
	errs []error,
) {

	if bidderRawResponse.StatusCode == http.StatusNoContent {
		fmt.Println("---------http.StatusNoContent")
		return nil, nil
	}

	if bidderRawResponse.StatusCode == http.StatusBadRequest {
		fmt.Println("---------http.StatusBadRequest")
		err := &errortypes.BadInput{
			Message: fmt.Sprintf(unexpectedStatusCodeMessage, bidderRawResponse.StatusCode),
		}
		return nil, []error{err}
	}

	if bidderRawResponse.StatusCode != http.StatusOK {
		fmt.Println("---------NOT http.StatusOK")
		err := &errortypes.BadServerResponse{
			Message: fmt.Sprintf(unexpectedStatusCodeMessage, bidderRawResponse.StatusCode),
		}
		return nil, []error{err}
	}

	fmt.Println("---------http.StatusOK")

	var openRTBBidderResponse openrtb.BidResponse
	// fmt.Println("---------openrtb.BidResponse\n", openrtb.BidResponse)
	if err := json.Unmarshal(bidderRawResponse.Body, &openRTBBidderResponse); err != nil {
		return nil, []error{err}
	}
	fmt.Println("---------openRTBBidderResponse\n", openRTBBidderResponse)

	bidsCapacity := len(openRTBBidderResponse.SeatBid[0].Bid) // e.g. SeatBid[0] is LD, Bid is an array?
	bidderResponse = adapters.NewBidderResponseWithBidsCapacity(bidsCapacity)

	var typedBid adapters.TypedBid
	for _, seatBid := range openRTBBidderResponse.SeatBid {
		for i := range seatBid.Bid {
			fmt.Println("---------seatBid.Bid[i]\n", seatBid.Bid[i])
			typedBid = adapters.TypedBid{Bid: &seatBid.Bid[i], BidType: "banner"}
			bidderResponse.Bids = append(bidderResponse.Bids, &typedBid)
		}
	}

	fmt.Println("---------bidderResponse\n", bidderResponse)
	return bidderResponse, nil
}


func NewLockerDomeBidder(endpoint string) *LockerDomeAdapter {
	return &LockerDomeAdapter{endpoint: endpoint}
}
