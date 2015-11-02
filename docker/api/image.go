package api

import (
	"encoding/json"
	"io/ioutil"
)

type Image struct {
	Created     int64
	Id          string
	ParentId    string
	RepoDigests []string
	RepoTags    []string
	Size        int64
	VirtualSize int64
}

type Images []Image

// ListImgs
func ListImgs() (imgs Images, err error) {
	res, err := client.GET("/images/json?")
	if err != nil {
		panic(err)
	}
	imgs = make(Images, 0)
	bs, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(bs, &imgs)
	return
}

// BuildImage ...
func BuildImage() {

}
