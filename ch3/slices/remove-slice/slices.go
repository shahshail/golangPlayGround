package main

import (
	"fmt"
)

//This programm will illustrate how to remove an element from any position from slice

func main() {
	s := []int{5, 6, 7, 8, 9, 10}
	fmt.Println(remove(s, 7))
}

func remove(slice []int, index int) []int {
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}
