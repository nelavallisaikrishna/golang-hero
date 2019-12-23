package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)


func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Eroor:", err)
		os.Exit(1)
	}
	// It will create byte with n(99999) number of elements
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	//or

	io.Copy(os.Stdout, resp.Body)
}
