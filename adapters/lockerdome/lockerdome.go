package lockerdome

import ( // cut down to the ones we absolutely need
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/prebid/prebid-server/pbs"

	"golang.org/x/net/context/ctxhttp"

	"github.com/mxmCherry/openrtb"
	"github.com/prebid/prebid-server/adapters"
	"github.com/prebid/prebid-server/errortypes"
	"github.com/prebid/prebid-server/openrtb_ext"
	"github.com/prebid/prebid-server/pbsmetrics"
)

type lockerdomeAdapter struct {
	http           *adapters.HTTPAdapter // do we need this?
	URI            string
	// iabCategoryMap map[string]string // do we need this?
}

// cookie handling is legacy???????????

type lockerdomeAdapterOptions struct {
  // todo - do we need this
}

lockerdomeParams struct { // check for redundant
  // find out what params we have access to. may not be the same as front end prebid
  adUnitCode           string          `json:"adUnitCode"` // ?
  adUnitId             string          `json:"adUnitId"`
  requestId            string          `json:"requestId"`
  sizes                string          `json:"sizes"` // ?
}

type lockerdomeImpExtLockerDome struct {
  // what's the diff between this and above???
}
