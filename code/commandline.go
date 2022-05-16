package main

import (
	"flag"
	"fmt"
	"os"
)

// global vars
var str string
var num int
var help bool

func main() {
	// go run commandline.go --str=Foo --num=8 --help=true filename
	num_args := len(os.Args)
	if num_args < 2 {
		fmt.Println(">> No args passed in")
	}

	// define flags
	flag.StringVar(&str, "str", "default value", "text description")
	flag.IntVar(&num, "num", 5, "text description")
	flag.BoolVar(&help, "help", false, "Display Help")

	// parse
	// 必须parse，否则都是默认值
	flag.Parse()

	// check if help was called explicitly
	if help {
		fmt.Println(">> Display help screen")
		os.Exit(1)
	}

	// See values assigned
	fmt.Println(">> String:", str)
	fmt.Println(">> Number:", num)

	// last command-line argument
	fmt.Println(">> Last Item:", os.Args[num_args-1])

	// the os.Args will include flags for example
	// go run command-line-args.go --str=Foo filename
	// os.Args[1] = "--str=Foo"

	// If you have flags and want just the arguments
	// then after calling flag.Parse() you can call
	// flag.Args which store only the args
	args := flag.Args()
	if len(args) > 0 {
		fmt.Println(">> Flag Arg:", args[0])
	}

}
