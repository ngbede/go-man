package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// this tells go to create an array of strings and add a bunch of functions designed to work with it
// the functions are called `receiver`
// A function with a receiver is like a method of a class
type deck []string

func mains() {
	os.Chdir("../..")
	deckOfCards := newDeck()
	// deckOfCards.print()

	deckOfCards[4:7].print() // slices in go
	// deckOfCards[32:40].print()
	// deckOfCards[:4].print()
	// deckOfCards[50:].print()

	fmt.Println("Dealing...")
	dealed, cardsLeft := deal(deckOfCards, 3)
	dealed.print()
	cardsLeft.print()

	c := color("Red")
	var da color = "joke"
	fmt.Println(da.describe("joker"))
	fmt.Println(c.describe("is an awesome color"))

	// type conversion
	jag := []byte("Hello World")
	fmt.Println(jag)

	dealed.saveToFile("lol")
	fmt.Println(deckOfCards.toString())

	fmt.Println(os.Getwd())
	readFile("lol.txt")
}

// variables of type deck gain access to the print method below by default
// define the func then the receiver (argName type) functionName() {...function body}
func (d deck) print() {
	for i, cards := range d {
		fmt.Println(i+1, cards)
	}
}

func (d deck) toString() string {
	stringArray := []string(d)
	s := strings.Join(stringArray, ",")
	return s
}

// creates and returns a list of playing cards i.e a deck
func newDeck() deck {
	cards := deck{} // empty deck of cards
	cardSuites := [4]string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Joker", "Queen", "King"}

	// replacing a value not being used with an underscore _, tells the go compiler to ignore it in error checks
	for _, i := range cardSuites {
		for _, j := range cardValues {
			var card string = j + " of " + i
			cards = append(cards, card)
		}
	}
	return cards
}

func deal(cards deck, handSize int) (deck, deck) {
	dealedCards := cards[:handSize]
	cardsLeft := cards[handSize:]
	return dealedCards, cardsLeft
}

type color string

func (c color) describe(description string) string {
	return string(c) + " " + description
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName+".txt", []byte(d.toString()), 0666)
}

func readFile(fileName string) deck {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1) // quits the program
	}
	str := string(bytes)
	strArray := strings.Split(str, ",")
	return deck(strArray)
}

func (d deck) shuffle() {
	randInt := time.Now().UnixNano()
	source := rand.NewSource(randInt)
	r := rand.New(source)

	for index := range d {
		// change seed value to make generation actually random
		newPosition := r.Intn(len(d) - 1)
		// swapping logic
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
