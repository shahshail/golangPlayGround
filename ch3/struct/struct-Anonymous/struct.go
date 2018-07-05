package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	circle Circle
	Spokes int
}

func main() {
	//This application may be clearer for it, but this code structure makes accessible the fields of Wheel more verbose
	var wheel Wheel
	wheel.circle.Center.X = 8
	wheel.circle.Center.Y = 8
	wheel.circle.Radius = 5

	fmt.Println(wheel)

}
