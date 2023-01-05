package main

import (
	"fmt"
)

func main() {
	numArray := [11]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := make([]int, 11)
	fmt.Println(n)
	// fmt.Println(numArray)

	for _, val := range numArray {
		if val%2 == 0 {
			fmt.Println(val, "is even")
		} else {
			fmt.Println(val, "is odd")
		}
	}

	// for i := range n {
	// 	if i%2 == 0 {
	// 		fmt.Printf("%v is even\n", i)
	// 	} else {
	// 		fmt.Printf("%v is odd\n", i)
	// 	}
	// }
}
