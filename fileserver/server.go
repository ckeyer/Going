package main

import (
	"fmt"
	"net/http"
)

func main() {
	h := http.FileServer(http.Dir("./"))
	fmt.Println("Listenning on :2222")
	if err := http.ListenAndServe(":2222", h); err != nil {
		panic(err)
	}
}
