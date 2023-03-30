package main

import "fmt"

func main() {
	numLetters := lettersOfWord("Sofia")
	fmt.Println(numLetters)
	printLetters("Sofia")
}

func lettersOfWord(word string) int {
	return len(word)
}

func printLetters(word string) {
	numLetters := lettersOfWord(word)
	for i := 0; i < numLetters; i++ {
		letra := word[i]
		fmt.Println(string(letra))
	}
}
