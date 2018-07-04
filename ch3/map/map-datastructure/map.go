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

	//Map can also be initialized by empty set
	age2 := map[string]int{
		"Shail": 26,
		"Lalu":  27,
	}
	fmt.Println(age2) // Equivellant to above map

	//To create Empty map
	emptyMap := map[string]int{}
	fmt.Println(emptyMap)

	//To Delete or Remove an item
	delete(ages, "Shail")
	fmt.Println(ages)

	//Map returns '0' when there is no key available.. it is fine if we dont have numeric value. But iof we have a numeric value then it is difficult to find either '0' means value or nil
	//We can check the status by
	age, ok := ages["bob"]           // we dont have bob as a key
	fmt.Printf("age : %d", age)      // 0
	fmt.Printf("Available : %b", ok) // returns false
}
