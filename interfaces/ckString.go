package main

import "fmt"

type CKString struct {
	Id   int
	Name string
}

func (this *CKString) String() string {
	return fmt.Sprint(this.Id) + ": " + this.Name
}

func main() {
	c := &CKString{Id: 12, Name: "nininin"}
	fmt.Println(c)
}
