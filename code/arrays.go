package main

import "fmt"

func main() {
	// The specification for slices are []T where T is the type.
	// for example []string is a set of strings,
	// or []int is a set of integers.
	var nums []int
	alphas := []string{"abc", "def", "ghi", "jkl"}

	nums = append(nums, 203)
	nums = append(nums, 303)
	fmt.Println(nums)
	alphas = append(alphas, "mno", "pqr", "stu")
	fmt.Println(alphas)

	fmt.Println("Length: ", len(alphas))
	fmt.Println(alphas[1])
	alpha2 := alphas[1:3]
	fmt.Println(alpha2)

	if elemExists("def", alpha2) {
		fmt.Println("Exists!")
	}
}

func elemExists(s string, a []string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}
