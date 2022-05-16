package main

import "fmt"

func main() {
	alphas := []string{"abc", "def", "ghi"}

	// standard for loop
	for i := 0; i < len(alphas); i++ {
		fmt.Printf("%d: %s \n", i, alphas[i])
	}

	// counting backwards
	for i := len(alphas) - 1; i >= 0; i-- {
		fmt.Printf("%d: %s \n", i, alphas[i])
	}

	for i, val := range alphas {
		fmt.Printf("%d: %s \n", i, val)
	}

	// while
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}

	// Infinite Loop
	// for {
	// 	fmt.Print(".")
	// }
}
