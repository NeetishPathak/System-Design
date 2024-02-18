package main

import "fmt"

func Initialize() {
	dict := Newnode()
	// dict.addWord("ANIME", 0, "")
	// dict.addWord("ANT", 0, "")
	// dict.addWord("ARMOR", 0, "")
	dict.addWord("APPLE", 0, "")
	// dict.addWord("AGILE", 0, "")
	// dict.addWord("AGE", 0, "")
	fmt.Println(dict)
	// fmt.Println("APPLE", dict.findWord("APPLE", 0, ""))
	fmt.Println("APPLO", dict.findWord("APPLO", 0, ""))
	// fmt.Println("AGILE", dict.findWord("AGILE", 0, ""))
	// fmt.Println("GAME", dict.findWord("GAME", 0, ""))
}

func main() {
	Initialize()
}
