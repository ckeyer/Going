package api

import (
	"testing"
)

// TestNewTlsClient ...
func TestNewTlsClient(t *testing.T) {
	return
	imgs, err := ListImgs()
	if err != nil {
		t.Error(err)
	}
	if imgs == nil {
		t.Skip("Images is nil")
	}
	for _, v := range imgs {
		log.Noticef("RepoTags: %s", v.RepoTags)
		log.Infof("Created: %s", v.Created)
		log.Infof("Id: %s", v.Id)
		log.Infof("VirtualSize: %s", v.VirtualSize)
	}
}
