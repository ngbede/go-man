package main

import (
	"fmt"
	"strconv"
)

// purpose of interfaces
// 1. Allows for code reuse between diff parts of the codebase

// a new type called bot that is accessible to all members of the code
// if any other type has a method of name getGreeting() which returns string, they automatically extend the bot interface functionality
// this sounds like inheritance in OOP to me
// they are used to define a method/function set
// think of it as a base class
// we can add multiple function definitions to the interface, any type extending the interface must implement all the base functions defined here..
// alongside the arg and return patterns
type bot interface {
	getGreeting() string
	greetName(string, int64) string
}

type englishBot struct{}

type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)

	fmt.Println(eb.greetName("Emmna", 1))
	fmt.Println(sb.getGreeting())
}

// as the receiver value is not being used, we can just pass the type to receier function
func (englishBot) getGreeting() string {
	return "Hello my friend"
}

func (spanishBot) getGreeting() string {
	return "Holla, mochas"
}

func (englishBot) greetName(name string, id int64) string {
	str := strconv.FormatInt(id, 10)
	return "Hello " + str + " " + name
}

// we'll refactor both functions to allow for reuse
// at this point we can pass any instance of a type with the getGreeting() method to the function below
// and it will execute successfully
func printGreeting(b bot) { // we pass the bot interface
	fmt.Println(b.getGreeting())
}

// Concrete types: are types we can create values out of e.g int, string, struct, string, custom types
// Interface types: are types that values/instances can't be created out of

// Extra notes
// 1. Interfaces are not generic types, no support for generic types in Go
// 2. Interfaces are implicit i.e no need to manually define what types extend off of the interface once its defined
// 3. Interfaces are a contract to help us manage types, if my custom types implementation of a func are broken interface won't help pick that up
// 4. Interfaces might be tought, requires experience in reading docs and writing custom interfaces

// We can specify inerfaces as a type inside of a struct
// We can also take multiple interfaces and assemble them into a single interface body e.g
type Reader interface{}
type Closer interface{}
type ReadCloser interface {
	Reader
	Closer
}
