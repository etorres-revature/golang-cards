package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("error occured - ", err)
		os.Exit(1)
	}

	fmt.Println("response:", resp)

	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println("response body:", string(bs))

}
