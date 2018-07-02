package main

import "fmt"

func mainPointer()  {
	// This is Pointer tutorial..
	var x int  = 1
	var p = &x
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)
	fmt.Println("THe address of local pointers are" , x , *p)
	fmt.Println(f())
	
	// Lets increment a number
	v := 1 
	incr(&v)

}

var p = f()

func f() *int {
	v := 1
	return &v // for each call, f returns a distinct value..
}

func incr(p *int) int {
	*p++ // increments what p points to; does not change p
	return *p
}