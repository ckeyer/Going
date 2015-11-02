package ini

import (
	"fmt"
	"testing"
)

func TestGetIni(t *testing.T) {
	kv, err := GetIniConfig("a.ini")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%#v\n", kv)
	k, ok := kv["asdf"]
	fmt.Printf("%#v, %#v\n", k, ok)
}
