package main

//This program will illustrate the simple map data structure in GO. Map is HashTable or Hashmap datastructure that contain key and value pair

import "fmt"

func main() {
	fmt.Println("HashTable-Map Example")
	ages := make(map[string]int) // Here key=string and value=int (map[k]v)
	ages["Shail"] = 26
	ages["Lalu"] = 27

	fmt.Println(ages)
	fmt.Println(ages["Shail"]) //26

}
