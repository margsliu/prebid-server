package lockerdome

import (
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

// LockerDome does not currently support syncing

func TestLockerDomeSyncer(t *testing.T) {
	temp := template.Must(template.New("sync-template").Parse("localhost"))
	syncer := NewLockerDomeSyncer(temp)
	syncInfo, err := syncer.GetUsersyncInfo("", "")
	assert.NoError(t, err)
	assert.Equal(t, "localhost", syncInfo.URL)
	assert.Equal(t, "redirect", syncInfo.Type) // "redirect" or "iframe"? prob. redirect. or, it doesn't really apply to us?
	assert.EqualValues(t, 0, syncer.GDPRVendorID())
	assert.Equal(t, false, syncInfo.SupportCORS)
}
