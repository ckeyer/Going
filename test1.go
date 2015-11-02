package main

import "fmt"
import "gopkg.in/mgo.v2/bson"

type BB map[string]int

func main() {

	var a struct {
		Name BB
	}
	a.Name = make(BB)
	a.Name["name"] = 12
	var b bson.ObjectId
	b = bson.NewObjectId()

	fmt.Printf("%#v\n", b.Hex())
	fmt.Printf("%#v", a)

	obtained := "compose.Revision{ID:\"V3\\x17sj\\xe70\\x00\\x0e\\x00\\x00+\", TemplateID:\"V3\\x17sj\\xe70\\x00\\x0e\\x00\\x00*\", ServiceGroup:compose.ServiceGroup{\"test_service\":(*compose.Service)(0xc820089800)}, Description:\"\", CreatedAt:time.Time{sec:63581785715, nsec:53000000, loc:(*time.Location)(0x118e100)}}"
	expected := "compose.Revision{ID:\"V3\\x17sj\\xe70\\x00\\x0e\\x00\\x00+\", TemplateID:\"V3\\x17sj\\xe70\\x00\\x0e\\x00\\x00*\", ServiceGroup:compose.ServiceGroup{\"test_service\":(*compose.Service)(0xc820089400)}, Description:\"\", CreatedAt:time.Time{sec:63581785715, nsec:53870817, loc:(*time.Location)(0x118e100)}}"
	fmt.Println(obtained == expected)
	fmt.Println(len(obtained))
	fmt.Println(len(expected))
	for i := 0; i < len(obtained) && i < len(expected); i++ {
		if obtained[i] == expected[i] {
			continue
		}
		fmt.Printf("%d,%s,%s\n", i, string(obtained[i]), expected[i])
	}
}
