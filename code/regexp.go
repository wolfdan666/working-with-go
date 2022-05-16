package main

import (
	"fmt"
	"regexp"
)

func main() {
	pattern := "[0-9]+"

	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex", err)
	}

	str := "The 12 monkeys ate 48 bananas"
	if re.MatchString(str) {
		fmt.Println("Yes, matched a number")
	} else {
		fmt.Println("No, no match")
	}

	result := re.FindString(str)
	fmt.Println("Number matched:", result)

	// specify -1 to return all.
	results := re.FindAllString(str, 2)
	for i, v := range results {
		fmt.Printf("Match %d: %s\n", i, v)
	}

	result = re.ReplaceAllString(str, "xx")
	fmt.Println("Result:", result)

	// Here is an example using the case insensitive flag i
	// ptn := `(?i)^t.`  // Match 0: To
	ptn := `(?i)t.` // Match 0: ToMatch 1: t Match 2: to
	str = "To be or not to be"

	re, err = regexp.Compile(ptn)
	if err != nil {
		fmt.Println("Error compiling regex", err)
	}

	// match string
	results = re.FindAllString(str, -1)
	for i, v := range results {
		fmt.Printf("Match %d: %s", i, v)
	}
	fmt.Println()

	// Submatches are what Go calls the matches that
	// are grabbed by (.*) inside of a regular expression.
	str1 := "Hello @world@ Match"
	sub_re, _ := regexp.Compile("@(.*)@")

	m := sub_re.FindStringSubmatch(str1)
	fmt.Println(m)
	// FindStringSubmatch returns [@world@ world]
	// so to just get the submatch you would use
	if len(m) > 1 {
		fmt.Println(m[1]) // submatch
	}

	// If you wanted to match brackets or other special characters
	// and try to just escape them like so \[(.*)\]
	// you will get an error for unknown escape sequence [

	// You need to double up the slashes or
	// a nicer solution is to use string literals
	// and wrap in ticks instead of quotes

	str2 := "A [word] here and [there]"
	esc_pattern := `\[(.*?)\]`
	esc_re, _ := regexp.Compile(esc_pattern)

	// this will only find the first
	fmt.Println(esc_re.FindStringSubmatch(str2))

	// use FindAll with second parameter for # of matches -1 = all
	fmt.Println(esc_re.FindAllStringSubmatch(str2, -1))
}
