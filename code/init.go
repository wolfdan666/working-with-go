package main

import "fmt"

var phrase string

func init() {
	phrase = "Init First!"
}

func main() {
	fmt.Println(phrase)
}
