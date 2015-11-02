package api

import (
	"encoding/json"
	"fmt"
)

type AA struct {
	User   string `json:"user,omitempty"`
	Sec    string `json:"sec,omitempty",`
	Key    string `json:"key,omitempty"`
	Values string `json:"values,omitempty"`
	Size   int    `json:"size,omitempty"`
}

func EncodeJson() {
	a := &AA{
		User:   "hero",
		Sec:    "anxciwe",
		Values: "lalalla",
		Size:   10,
	}
	bs, _ := json.Marshal(a)
	fmt.Printf("%s\n", bs)
}
