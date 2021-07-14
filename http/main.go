package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("error occured - ", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	return 1, nil
}
