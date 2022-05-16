package main

import "fmt"

func main() {
	num := 4
	if num > 3 {
		fmt.Println("Many")
	} else if num == 2 {
		fmt.Println("Two")
	} else {
		fmt.Printf("Others")
	}

	switch {
	case num == 1:
		fmt.Println("One")
	case num == 2:
		fmt.Println("Two")
	case num > 2:
		fmt.Println("Many")
	default:
		fmt.Println("Less than 1")
	}
}
