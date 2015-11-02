package api

import (
	"testing"
)

func TestVersion(t *testing.T) {
	v, err := GetVersion()
	if err != nil {
		t.Error(err)
	}
	log.Noticef("%#v\n", v)
}

func TestPing(t *testing.T) {
	ok, err := Ping()
	if err != nil {
		t.Error(err)

	}
	log.Notice("Ping is ", ok)
}
