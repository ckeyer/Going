package main

import (
	"encoding/json"
	"fmt"
)

var str = `{
"name":"ckeyer",
"age":23,
"marriged":true,
"profile":{
"email":"ckeyer.com"
}
}
`

type Person struct {
	Name    string  `json:"name""`
	Age     int     `json:"age"`
	Mar     bool    `json:"marriged"`
	Profile Profile `json:"profile"`
}
type Profile struct {
	Email string `json:"email"`
}

func main() {
	a := new(Person)
	json.Unmarshal([]byte(str), a)
	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", a.Profile)
}
