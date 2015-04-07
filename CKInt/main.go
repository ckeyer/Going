package main

import (
	"fmt"
)

type CKBaseInt int8

const (
	CK_BASE_LEN = 8
)

type CKInt struct {
	V      []CKBaseInt
	Length int
}

func NewCKInt(s string) *CKInt{
	
}

func (this *CKInt) ToString() string{
	sout:=""
	for _,v range this.V{
		sout += fmt.Sprint(v)
	}
	return sout
}

func main() {

}
