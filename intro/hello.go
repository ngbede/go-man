package main

import (
	"fmt"
	"math"
)

func main() {
	age := 22
	tall := true
	name := "Emmanuel Sule"

	fmt.Println("I am ", age, " years old")
	fmt.Println(tall, name)
	fmt.Println("Hello World!")

	arithmetic := math.Sqrt(float64(age))
	fmt.Println(arithmetic)

	for i := 0; i < 100; i++ {
		fmt.Println("Loop", i)
	}
}

// go build: compiles the source code into an executable file
// go run: compiles and executes one or more go files

// a `package` is a collection of common src code files. Same as project or workspace.
// package can have many related files inside of it.
// Every file must declare the package they belong to i.e `package main` at the very 1st line
// Types of packages: executables (generates a file we can run) & reusable (code used as helpers) => put reusable logic e.g code dependency, libraries
// Difference between the two above: the name of the package used determines if its an executable or reusable type package.
// `package main` => used for executable type packages while `package anyOtherName` => used for reusable package names
// `package main` must always have a function called `main`
