package main

import "fmt"

// how we normally declare const vars
const ALIVE bool = true
const PLANETS = 9

var person string // = "Phil"

// var lol := 23 we can't initialize a variable outside a function using this syntax

func main() {
	// Variables can be initialized outside of a function, but cannot be assigned a variable.
	var card string = "Ace of Spades"
	name := "Embid" // type inference syntax, type inference can only be used in variable declaration, not after
	fmt.Println(card, name, person)

	person = "Joseph" // NOT name := "Joseph"
	fmt.Println(person)
	// Types: string, bool, int, float32, float64

	// array definitions
	arr := [20]string{"apples", "oranges", "beans"} // \syntax 1
	var schools [6]string                           // syntax 2

	schools[3] = "Omega Turners"
	arr[19] = "Fish"
	fmt.Println(arr)
	fmt.Println(schools)

	fmt.Println(newCard())
	var addMe int = add(4, 5)
	fmt.Println(addMe)

	// GO has 2 types of arrays, each array definition must have its type specified
	// Array: fixed length list of items e.g var arr [len]type or arr := [len]type{...items}
	// Slice: an array that can shrink or grow in size

	cards := []string{newCard(), msg("lol")}
	cards = append(cards, "logical fush") // add new item to cards slice and reassign
	fmt.Println(cards)

	// loop over items in array
	for i, v := range cards {
		fmt.Println(i+1, v)
	}

	lengh := len(cards)
	fmt.Println(lengh)

	// fmt.Println("ha ha")
	// nc := newDeck()
	// nc.shuffle()
	// nc.print()
}

func newCard() string {
	return "Five of Diamonds"
}

func add(a int, b int) int {
	return a + b
}

func msg(name string) string {
	return "Hello " + name
}
