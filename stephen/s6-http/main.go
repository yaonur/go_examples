package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//bs := make([]byte, 99999)
	//read, err := resp.Body.Read(bs)
	//if err != nil && err != io.EOF {
	//	fmt.Println("Error:", err)
	//}
	//fmt.Println(read)
	//fmt.Println(string(bs))
	lw := logWriter{}
	_, err = io.Copy(lw, resp.Body)
	if err != nil && err != io.EOF {
		fmt.Println("Error:", err)
	}

}
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
