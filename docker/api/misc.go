package api

import (
	"encoding/json"
	"io/ioutil"
)

type Version struct {
	DockerVersion string `json:"Version"`
	Os            string
	KernelVersion string
	GoVersion     string
	GitCommit     string
	Arch          string
	ApiVersion    string
}

func GetVersion() (v *Version, err error) {
	res, err := client.GET("/version")
	if err != nil {
		return
	}
	bs, _ := ioutil.ReadAll(res.Body)
	v = new(Version)

	err = json.Unmarshal(bs, v)
	return
}

// Ping ...
func Ping() (ok bool, err error) {
	res, err := client.GET("/_ping")
	if res.StatusCode == 200 {
		ok = true
		return
	}
	return
}
