package main

import "fmt"
import "os/exec"

func main() {
	script := "`cd dir && touch aaa.txt`"
	out, err := exec.Command("bash", "-c", script).Output()
	if err != nil {
		fmt.Errorf("WebHook Error. %s", err)
		return
	}

	out, err = exec.Command("bash", "-c", "pwd").Output()
	if err != nil {
		fmt.Errorf("WebHook Error. %s", err)
		return
	}
	fmt.Printf("%s", out)
}
