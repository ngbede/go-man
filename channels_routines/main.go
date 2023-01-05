package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	sites := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com/",
		"https://github.com/ngbede",
		"https://golang.org",
		// "http://127.0.0.1:3000/api/v1/book/search/sxVHDwAAQBAJ",
	}

	ch := make(chan string)

	for _, s := range sites {
		go checkLink(s, ch)
	}
	// fmt.Println(<-ch) // this is a blocking call as go is waiting for a routine to send in a message

	// for i := 0; i < len(sites); i++ {}
	for l := range ch { // infinite while loop
		// fmt.Println(<-ch) // still a blocking call that blocks the for loop from proceeding until a message is received in the channel
		// go checkLink(l, ch) // repeating the routine call at this step coninuously

		// lets use a function literal to do the pausing of code before running the checlist function
		// we never try to reference the same variable inside of 2 diff routines
		go func(link string) {
			time.Sleep(time.Second * 5) // pause for 5 secs
			checkLink(link, ch)         // we should pass the link to this func as an arguement from the main routine
		}(l) // we are passing arguements to the anonmous function here
	}
}

func checkLink(s string, ch chan string) {
	_, err := http.Get(s) // this is a blocking call
	if err != nil {
		fmt.Println(s, "seems to be down")
		ch <- s
		return
	}
	ch <- s
	fmt.Println(s, "site is up")
}

// Go executes go code via a GO routine (engine that executes the code)
// To launch a new go routine to execute our code, we simply add the go keyword in front of the peice of code we want to run in parallel
// e.g go checkLink("https://youtube.com")
// Subsequent go routines get spawned from the main go routine
// we only use the go keyword in fron of function calls
// By default go routines execute on a single cpu core concurrently, we can override them and execute the routines onmultple cores (parallelism)
// THe main rouine is he sole entity responsible for controlling when our program quits/exits

// Channels: are used to communicate between diff running go routines, they are the only we have to communicate between go routines
// Info passed to a channel must of be of the same type

// channel syntax
// `channel <- 5` => sends the value 5 into the channel
// `myVar <- channel` => waits for a value sent from the channel and assigns it to a variable called myVar
// `<- channel` => simply waits for a value to be sent from a channel when we get one
// receving messages from a channel is a blocking call

// always pass variables from the main routine to the child routine via arguements or using channels
// turns out if you declare a channel variable and never use it in your go program, it just hangs
