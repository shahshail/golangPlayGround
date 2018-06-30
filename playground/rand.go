package main

import ("fmt"
		"math/rand")

var name string = "Shail" 

func main()  {
	name := "shah"
	fmt.Println("Hello from Go and the random number is %i" , rand.Intn(100))	
	fmt.Printf("Hello my name is  %s", name)
}