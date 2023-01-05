package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWritter struct{}

func main() {
	res, err := http.Get("http://127.0.0.1:3000/api/v1/book/search/sxVHDwAAQBAJ")

	if err != nil {
		fmt.Println(err.Error())
	}

	// Reader Interface: takes in multiple source of inputs and outputs a byte slice []byte (output data anyone can work with)
	// data := make([]byte, 3999)
	// the Read function continuosly sticks data into the byte slice until the slice is full, so it stops reading data into it once the byte slice is full
	// bytesRead, _ := res.Body.Read(data)
	// fmt.Println(bytesRead)
	// fmt.Println(string(data)) // log the byte data by passing it as a string

	// 2nd way of loggin out the data via the writer interface
	// pass our custom logWritter which now implements the writer interface i.e the Write func below
	lw := logWritter{}
	io.Copy(lw, res.Body) // because we pass interface that implements the writer interface, it doesn;t guarantee that it would work
}

// this logWritter is now implementing the writter interface
func (l logWritter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("just wrote this amount of bytes: ", len(bs))
	return len(bs), nil
}
