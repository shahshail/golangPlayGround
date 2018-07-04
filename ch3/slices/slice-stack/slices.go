package main

import (
	"fmt"
)

//This programm will create a custome stack using slice DataStructure

func main() {
	stack := make([]int, 10)

	stack = append(stack, 10)    //push
	top := stack[len(stack)-1]   // peek
	stack = stack[:len(stack)-1] // pop

	fmt.Println(top)
}
