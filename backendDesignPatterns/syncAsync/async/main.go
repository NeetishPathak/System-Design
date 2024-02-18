package main

import (
	"fmt"
	"os"
	"sync"
)

// dowork1
// readFile() <<-- non blocking
// dowork2

var filename string = "../dat.txt"

func dowork(i int) {
	fmt.Println("Doing work", i)
}

func readFile(wg *sync.WaitGroup) {
	defer wg.Done()
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
	var wg sync.WaitGroup
	dowork(1)
	wg.Add(1)
	go readFile(&wg)
	dowork(2)
	wg.Wait()
}
