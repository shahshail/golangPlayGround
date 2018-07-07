package main

import (
	"fmt"
)

func main() {
	fmt.Println(name("Shail"), surname("Shah"))
}

func name(name string) string {
	return name
}

func surname(surname string) string {
	return name(surname)
}
