package main

import (
	"encoding/json"
	"fmt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	db_name = "ckeyer"
	db_url  = "localhost:27017"
)

type (
	Blog struct {
		Id      string `bson:"_id"`
		BlogId  int64
		Title   string
		Content string
		Profile *Profile
		Author  []*User
	}
	Profile struct {
		Email string
		Addr  string
	}
	User struct {
		Name string
		Age  int
		Mar  bool
	}
)

func main() {
	_ = bson.M{"sdf": 1}
	_ = &json.Decoder{}
	fmt.Println("Start")
	var b1 Blog
	session, err := mgo.Dial(db_url)
	if err != nil {
		fmt.Println(err)
		return
	}
	c := session.DB(db_name).C("blog")
	err = c.Find(bson.M{"_id": "config"}).One(&b1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(b1)
	fmt.Println("###########")
	fmt.Println(*b1.Profile)
	fmt.Println("###########")
	for _, v := range b1.Author {
		fmt.Println(*v)
	}

	bs, err := json.Marshal(b1)
	fmt.Println(string(bs))
	fmt.Println("Over")
}

func insert() {
	b1 := &Blog{
		Id:      "config",
		BlogId:  12,
		Title:   "namestring",
		Content: "contentstring",
		Profile: &Profile{
			Email: "emailstring",
			Addr:  "addrstring",
		},
		Author: []*User{
			&User{
				Name: "adfsa",
				Age:  123,
				Mar:  true,
			},
			&User{
				Name: "gwl",
				Age:  123,
				Mar:  false,
			},
		},
	}

	session, err := mgo.Dial(db_url)
	if err != nil {
		fmt.Println(err)
		return
	}
	c := session.DB(db_name).C("blog")
	if err = c.Insert(b1); err != nil {
		fmt.Println(err)
		return
	}

}
