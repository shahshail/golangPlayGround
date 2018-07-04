package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("HashTable-Map Example")
	ages := make(map[string]int) // Here key=string and value=int (map[k]v)
	ages["Shail"] = 26
	ages["Lalu"] = 27

	fmt.Println(ages)

	var names []string
	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("$s\t%d\n", name, ages[name])
	}
}
