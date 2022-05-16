package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	path := "./testDir"

	// 类似于 chroot path && find /
	filepath.Walk(path, func(fn string, fi os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Walker Error: ", err)
			return err
		}

		if fi.IsDir() {
			fmt.Println("Directory: ", fn)
			// 这个需要全路径名才能skip, 否则结果输出File:  testDir/skipme
			if fi.Name() == "skipme" {
				return filepath.SkipDir
			}
		} else {
			fmt.Println("File: ", fn)
		}
		return nil
	})

	dir := "subDir"
	pathdir := filepath.Join(path, dir)
	if _, err := os.Stat(pathdir); os.IsNotExist(err) {
		mdErr := os.Mkdir(pathdir, 0755)
		if mdErr != nil {
			fmt.Println("Error making directory", mdErr)
		}
	}

	if usr, err := user.Current(); nil != err {
		fmt.Println("can't get user home")
	} else {
		fmt.Println("User Home Directory:", usr.HomeDir)
	}
}
