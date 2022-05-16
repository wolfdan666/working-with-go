package main

import "fmt"

// define Dog struct
type Dog struct {
	Name  string
	Color string
}

func main() {
	Spot := Dog{Name: "Spot", Color: "brown"}
	var Rover Dog
	Rover.Name = "Rover"
	Rover.Color = "light tan with dark patches"
	fmt.Println(Spot, Rover)
}
