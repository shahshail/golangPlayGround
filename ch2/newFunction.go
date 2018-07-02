// An another way to create/declare a variable in go is to create a 'new function'
package main

import "fmt"

func main() {
	p  := new(int) // p, of type *int points to unnamed int variable
	*p = 2

	fmt.Println(p)
	fmt.Println(*p)
}
