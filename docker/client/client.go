package client

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"io/ioutil"
	"net/http"
)

var (
	client   *http.Client
	endpoint = "https://192.168.59.103:2376"
)

// init
func init() {
	path := "../auth"
	ca := fmt.Sprintf("%s/ca.pem", path)
	cert := fmt.Sprintf("%s/cert.pem", path)
	key := fmt.Sprintf("%s/key.pem", path)

	c, err := docker.NewTLSClient(endpoint, cert, key, ca)
	if err != nil {
		panic(err)
	}
	client = c.HTTPClient
}

// ListImgs ...
func ListImgs() error {
	res, err := client.Get(endpoint + "/images/json?")
	if err != nil {
		panic(err)
	}
	bs, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("%s\n", bs)
	return nil
}
