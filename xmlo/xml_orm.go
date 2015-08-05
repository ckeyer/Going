package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type XmlOrm struct {
	FilePath string
	Object   interface{}
}

type User struct {
	Id        int
	Name      string
	IsMarried bool
	Age       int
	Profile   *Profile
}

type Profile struct {
	Email    string
	PhoneNum string
	Pic      []byte
}

//
func (o *XmlOrm) Insert() (n int, err error) {
	f, err := os.OpenFile(o.FilePath, os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		return
	}
	defer f.Close()
	bs, err := xml.Marshal(o.Object)
	if err != nil {
		return
	}
	n, err = f.WriteString(xml.Header)
	if err != nil {
		return
	}
	n2, err := f.Write(bs)
	n += n2
	return
}

//
func (o *XmlOrm) Read() (err error) {
	f, err := os.OpenFile(o.FilePath, os.O_RDONLY, 0664)
	if err != nil {
		return
	}
	r := xml.NewDecoder(f)
	token, err := r.RawToken()
	if err != nil {
		return
	}
	fmt.Printf("Token: %s\n", token)
	var u User
	err = r.Decode(&u)
	if err != nil {
		return
	}
	fmt.Printf("%#v\n", u)
	return
}

//
func main() {
	u := &User{
		Id:        111,
		Name:      "ckeyer",
		IsMarried: false,
		Age:       23,
		Profile: &Profile{
			Email: "me@ckeyer.com</Email>",
			Pic:   []byte("hello world"),
		},
	}
	fmt.Println(xml.Header)

	u_bs, err := xml.Marshal(u)
	if err != nil {
		return
	}
	fmt.Printf("%s\n", u_bs)
	var me User
	err = xml.Unmarshal(u_bs, &me)
	if err != nil {
		return
	}
	xo := &XmlOrm{
		FilePath: "test.xml",
		Object:   u,
	}
	// n, err := xo.Insert()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println(me.Profile.Email, n)
	err = xo.Read()
	if err != nil {
		fmt.Println(err.Error())
	}

}
