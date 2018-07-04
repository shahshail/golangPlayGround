package main

import (
	"fmt"
)

func main() {
	mySlice := []string{"a", "b", "c", "d", "e"}
	fmt.Println(mySlice)      //[a,b,c,d,e]
	fmt.Println(mySlice[2:4]) //[c,d]

	// Okay so lets actually create a slice

	newSlice := make([]int, 0, 5)

	fmt.Println("---------------------------")
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	fmt.Println("---------------------------")

	for i := 0; i < 80; i++ {
		newSlice = append(newSlice, i)
		fmt.Println("Len:", len(newSlice))
	}

}
