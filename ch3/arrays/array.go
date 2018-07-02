//Arrays are nothing but a fixed-length sequence of zero and more elements of a perticular type.
package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main () {
	var a [3]int //it creates an array of integer with size 3
	fmt.Println(a) // prints an entire array
	
	fmt.Println(len(a)) // 3

	var r [3]int = [3]int{1,2}
	r[len(r)-1] = 4 //adds num 4 at last index position
	fmt.Println(r)

	// We can use ellipsis to determine length of an array at initialization time.
	q := [...]int{1,2,3} //length of an array is 3
	fmt.Println(q) //[1,2,3]

	currency := [...]Currency{USD,EUR,GBP,RMB}
	fmt.Println(currency)
}