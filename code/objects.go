package main

import "fmt"

type Dog struct {
	Name  string
	Color string
}

func (d Dog) Call() {
	fmt.Printf("Here comes a %s dog, %s.\n", d.Color, d.Name)
}

type MyInt int

func (i MyInt) Double() MyInt {
	return i * 2
}

func main() {

	// create instance of dog
	Spot := Dog{Name: "Spot", Color: "brown"}

	// call object method
	Spot.Call()

	x := MyInt(3)
	fmt.Println(x.Double())
}
