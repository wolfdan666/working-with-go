package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	filename := "./rabbits.txt"

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("Error reading file", filename)
		return
	}

	// The content is returned as a []byte and not a string.
	// You need to cast to a string to use as such,
	// for example to display.
	// Use string() to cast a []byte to string type.
	fmt.Println(string(content))

	// 其实这个不是非常完美，因为2行内容，会输出3行结果...
	lines := strings.Split(string(content), "\n")

	for i, v := range lines {
		fmt.Println(i, v)
	}

	if _, err := os.Stat("junk.foo"); os.IsNotExist(err) {
		fmt.Println(">>>")
		fmt.Println("File: junk.foo does not exist")
	}

	// The way to deal with large file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	// write file
	outfile := "output.txt"
	// 最终文件竟然是777...
	err = ioutil.WriteFile(outfile, content, 0644)
	if err != nil {
		fmt.Println("Error writing file: ", err)
	} else {
		fmt.Println(">>>")
		fmt.Println("Created: ", outfile)
	}

	af, err := os.OpenFile(outfile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error appending to file:", err)
	}
	defer af.Close()
	if _, err = af.WriteString("Appending this text"); err != nil {
		fmt.Println("Error writing to file:", err)
	}

	// Use the filepath package for working
	// with cross-platform paths properly
	fmt.Println(filepath.Join("a", "b", "file.ext"))
}
