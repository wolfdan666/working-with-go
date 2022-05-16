package main

import (
	"fmt"
	"math"
)

func Say(s string) string {
	phrase := "Hello " + s
	return phrase
}

func Say2(s string) (phrase string) {
	phrase = "Hello " + s
	return
}

func Divide(x, y float64) (float64, float64) {
	q := math.Trunc(x / y)
	r := math.Mod(x, y)
	return q, r
}

func Divide2(x, y float64) (q, r float64) {
	q = math.Trunc(x / y)
	r = math.Mod(x, y)
	return
}

func Sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(Say("world"))
	fmt.Println(Say2("world2"))
	fmt.Println(Divide(10, 3))
	fmt.Println(Divide2(11, 3))

	sum := Sum(1, 3, 5, 7)
	fmt.Println(sum)

	nums := []int{1, 2, 3, 4, 5}
	sum = Sum(nums...)
	fmt.Println(sum)
}
