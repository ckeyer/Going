package main

import (
	"fmt"

	"github.com/fsouza/go-dockerclient"
)

var (
	client *docker.Client
)

// init ...
func init() {
	NewClient()
}

// NewClient 新建连接
func NewClient() {
	endpoint := "https://192.168.59.103:2376"

	path := "auth"
	ca := fmt.Sprintf("%s/ca.pem", path)
	cert := fmt.Sprintf("%s/cert.pem", path)
	key := fmt.Sprintf("%s/key.pem", path)

	var err error
	client, err = docker.NewTLSClient(endpoint, cert, key, ca)
	if err != nil {
		panic(err)
	}
}

// 列出所有镜像
func ListImgs() {
	imgs, err := client.ListIMages(docker.ListImagesOptions{All: false})
	if err != nil {
		fmt.Printf("%s \n", err)
		return
	}

	for _, img := range imgs {
		fmt.Println("ID: ", img.ID)
		fmt.Println("RepoTags: ", img.RepoTags)
		fmt.Println("Created: ", img.Created)
		fmt.Println("Size: ", img.Size)
		fmt.Println("VirtualSize: ", img.VirtualSize)
		fmt.Println("ParentId: ", img.ParentID)
		fmt.Println()
	}
}

// CreateImg 创建新的镜像
func CreateImg() error {
	var redis docker.BuildImageOptions
	redis.Pull = true
	redis.
		client.BuildImage(arg0)
}

func main() {
	ListImgs()
}
