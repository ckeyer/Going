package main

import "fmt"
import "time"
import "encoding/json"

type jsonTime struct {
	t time.Time
	f string
}

func (j jsonTime) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("%d-%02d-%02d %02d:%02d", j.t.Year(), j.t.Month(), j.t.Day(), j.t.Hour(), j.t.Minute())), nil
}

func main() {
	x := map[string]interface{}{
		"foo": jsonTime{t: time.Now(), f: time.Kitchen},
		"bar": "baz",
	}
	data, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data)
}
