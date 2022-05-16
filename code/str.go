package main

import (
	"fmt"     // for standard output
	"strings" // for manipulating strings
)

func main() {

	// create a string variable
	str := "HI, I'M UPPER CASE"

	// convert to lower case
	lower := strings.ToLower(str)

	// output to show its really lower case
	fmt.Println(lower)

	// check if string contains another string
	if strings.Contains(lower, "case") {
		fmt.Println("Yes, exists!")
	}

	sentence := "I'm a sentence made up of words"
	words := strings.Split(sentence, " ")
	fmt.Printf("%v \n", words)

	sentence = "I'm    a     sentence made up of words"
	fields := strings.Fields(sentence)
	fmt.Printf("%v \n", fields)

	lines := []string{"one", "two", "three"}
	// 这里网页上的string少写了一个s，把strings写成了string
	str = strings.Join(lines, ",")
	fmt.Println(str)

	str = "The blue whale loves blue fish."
	newstr := strings.Replace(str, "blue", "red", 1)
	fmt.Println(newstr)

	// go1.12 the same as strings.ReplaceAll
	newstr2 := strings.Replace(str, "blue", "red", -1)
	fmt.Println(newstr2)

	newstr3 := strings.ReplaceAll(str, "blue", "red")
	fmt.Println(newstr3)

	path := "/home/mkaz"
	isAbsolute := strings.HasPrefix(path, "/")
	fmt.Println(isAbsolute)

	dir := "/home/mkaz/"
	// 这里网页错写成了 path
	hasTrailingSlash := strings.HasSuffix(dir, "/")
	fmt.Println(hasTrailingSlash)
}
