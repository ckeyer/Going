package main

import (
	"fmt"
	"net"
	"net/http"
)

// Index ...
func Index(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Host)
	w.Write([]byte("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC2UUS1Hd9oBo+FvKxNB7SVbEqrs123QlYgoS+oCzvAGpoEqfeX4E2qGrZVhaY+NlAtIuo1+Ex3/prGRrby4zO61sGRtOHP23QBr16GnQpMBYApylO2rerS9iBXcurBY9CLKIoU1EYDCY07rKsyTAcSam9CCenE16lKzif/QwntfI3qrVs3FYv7+bDM+iabZ1UxGrNKIh5c5hVIAXZeB3jM0uuBXALU9Y6gTIIQWqZ+QGlNwBxs1pPpfdiRHRZ7m3xfMK7bGgFkayQUmWdkPGgGObLuf3C4K28TjJ6eE6wj5vfDGCpHXB0AyzdnuA7cp+Bkrk7IyB8fYsPZnXuLFdU1 root@Ckeyer\n"))
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		w.Write([]byte("\n" + addr.String()))
	}
}

func main() {
	fmt.Println("Start...")
	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("/a", Index)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
