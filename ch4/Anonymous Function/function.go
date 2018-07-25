package main

type User struct {
	id      int
	name    string
	surname string
}

func main() {

	getName()
}

func getName() func() string {
	return func() string { return "Shah Shail" }
}
