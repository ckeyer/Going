package ini

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// GetValue ...
func GetIniConfig(file string) (s map[string]string, err error) {
	f, err := os.OpenFile(file, os.O_RDONLY, 0600)
	if err != nil {
		return
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	s = make(map[string]string)
	for {
		bs, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Printf("%s\n", bs)
		kv := strings.Split(string(bs), "=")
		fmt.Printf("%#v\n", kv)
		if len(kv) == 2 {
			s[kv[0]] = kv[1]
		}
	}

	return
}
