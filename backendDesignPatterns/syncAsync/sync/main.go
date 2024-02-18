package main

import (
	"fmt"
	"os"
)

// dowork1
// readFile() <<-- blocks
// dowork2

var filename string = "../dat.txt"

func dowork1() {
	fmt.Println("Doing work1")
}

func dowork2() {
	fmt.Println("Doing work2")
}

func readFile() {
	data, err := os.ReadFile(filename)
	check(err)
	fmt.Println(string(data))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dowork1()
	readFile()
	dowork2()
}
