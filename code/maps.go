package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["c"] = "Cyan"
	m["y"] = "Yellow"
	m["m"] = "Magenta"
	m["k"] = "Black"

	var m2 = map[string]string{
		"c": "Cyan",
		"y": "Yellow",
		"m": "Magenta",
		"k": "Black",
	}

	for k, v := range m2 {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}

	c, ok := m["p"]
	if ok {
		fmt.Println(c)
	}
}
