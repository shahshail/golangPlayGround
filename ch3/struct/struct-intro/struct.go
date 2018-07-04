package main

import (
	"fmt"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	Position  string
	Salary    int
	ManagerId int
}

func main() {
	var employee1 Employee = Employee{01, "Shail", "NA", "Software Engineer - 1", 3700, 241}
	var employee2 Employee = Employee{01, "Ajay", "NA", "Software Engineer - 2", 13700, 341}

	fmt.Println(employee1.Name)
	fmt.Println(employee2.Name)

	employee1.Name = "Shail Shah"

	fmt.Println(employee1.Name) // Shail Shah

}
