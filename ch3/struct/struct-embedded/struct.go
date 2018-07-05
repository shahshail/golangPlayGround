package main

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var wheel Wheel
	wheel.X = 8      //wheel.Circle.Point.X
	wheel.Y = 8      //wheel.Circle.Point.X
	wheel.Radius = 5 //wheel.Circle.Radius
	wheel.Spokes = 20
}
