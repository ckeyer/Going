package api

import (
	"fmt"
	logpkg "github.com/ckeyer/go-log"
	"github.com/fsouza/go-dockerclient"
)

var (
	log      = logpkg.GetDefaultLogger()
	client   *Client
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
	client = NewClient(c.HTTPClient)
}
