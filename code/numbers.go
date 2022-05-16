package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 网页里的例子错误有点多，问题不大
	var str string
	var ans int
	str = "42"
	ans, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error convert string to integer", err)
		return
	}

	fmt.Println(ans)

	str = strconv.Itoa(369)
	fmt.Println(str)

	i := 43
	y := 43.0
	x := float64(i)
	fmt.Println(i, y, x)

	pi := "3.14159"
	pi_x, err := strconv.ParseFloat(pi, 64)
	if err != nil {
		fmt.Println("Error convert string to float", err)
		return
	}

	// +3.141590e+000#
	print(pi_x)
	fmt.Println("")
	fmt.Println(pi_x)

	fmt.Printf("Variable x has type %T and value %v", x, x)
	/*
		这个最后的#很奇怪
		42
		369
		43 43 43
		+3.141590e+000
		3.14159
		Variable x has type float64 and value 43#
	*/
}
