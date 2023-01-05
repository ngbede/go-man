package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	fn := args[len(args)-1]
	fmt.Println(fn)

	file, err := os.Open(fn)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	io.Copy(os.Stdout, file) // log contents to console
}
