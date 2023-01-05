package main

import "fmt"

type fake struct {
	id   string
	data string
}

// maps is simply a collection of key-value pairs. It is a reference type
// keys and values are statically typed i.e they must all be of the same type

func main() {
	// init 1
	colors := map[string]string{
		"red":    "#ff1234",
		"orange": "#ff9032",
		"yellow": "#ff3902",
		"green":  "#ff3920",
	}
	fmt.Println(colors["red"])

	// init 2
	var newMap map[int]bool // empty map
	fmt.Println(newMap)

	// init 3
	map2 := make(map[string]string)

	colors["blue"] = "ff7548" // add value to the map
	fmt.Println(colors)

	map2["joker"] = "batmans"
	fmt.Println(map2)

	// delete item from a map
	delete(colors, "red")
	fmt.Println(colors)

	// loop over map to get key, value pairs
	keys := make([]string, 0, len(colors))
	values := make([]string, 0, len(colors))
	for k, v := range colors {
		fmt.Println(v)
		keys = append(keys, k)
		values = append(values, v)
	}
	fmt.Println(keys, values)

	// nesting maps with structs
	dummy := map[string]fake{
		"hello": {
			id:   "1",
			data: "hello",
		},
	}
	fmt.Println(dummy)
}
