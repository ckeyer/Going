package main

import (
	"fmt"
	"strconv"
)

const (
	CK_BASE_LEN = 8
)

type CKInt struct {
	V      []byte
	Length int
}

func NewCKInt(s string) *CKInt {
	//bs := make([]byte, len(s))
	bs := []byte(s)
	for {
		if bs[0] == '0' {
			bs = bs[1:]
		} else {
			break
		}
	}
	for i := 0; i < len(bs); i++ {
		b, _ := strconv.Atoi(string(bs[i]))
		bs[i] = byte(b)
	}
	for i := 0; i < len(bs)/2; i++ {
		bs[i], bs[len(bs)-1-i] = bs[len(bs)-1-i], bs[i]
	}
	return &CKInt{V: bs, Length: len(bs)}
}

func (this *CKInt) ToString() string {
	bs := make([]byte, len(this.V))
	copy(bs, this.V)
	for i := 0; i < len(this.V)/2; i++ {
		bs[i], bs[len(this.V)-1-i] = bs[len(this.V)-1-i], bs[i]
	}
	s := ""
	for _, v := range bs {
		s += fmt.Sprint(v)
	}
	return s
}

func (a *CKInt) Sum(b *CKInt) *CKInt {
	length := a.Length
	lastInt := func(bs []byte) byte {
		return bs[len(bs)-1]
	}
	if length < b.Length {
		length = b.Length
	} else if length == b.Length {
		if lastInt(b.V)+lastInt(a.V) >= 10 {
			length++
		}
	}
	cbs := make([]byte, length)
	i := 0
	for ; i < a.Length && i < b.Length; i++ {
		tmp := a.V[i] + b.V[i]
		if tmp >= 10 {
			cbs[i+1]++
			tmp -= 10
		}
		cbs[i] += tmp
	}
	lastbytes := func(bss []byte) {
		if len(bss) > i {
			for ; i < len(bss); i++ {
				cbs[i] += bss[i]
			}
		}
	}
	lastbytes(a.V)
	lastbytes(b.V)

	return &CKInt{V: cbs, Length: len(cbs)}
}

func main() {
	a := NewCKInt("990000000000567")
	b := NewCKInt("9234567")
	c := a.Sum(b)
	fmt.Printf("%s + %s = %s\n", a.ToString(), b.ToString(), c.ToString())
}
