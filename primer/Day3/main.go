package main

import "fmt"

// 1. Creates a map to store the number of occurrences of each letter in a given string.
func frequencyCounter(word string) {
	counter := make(map[string]int)
	wordLength := len(word)
	for i := 0; i < wordLength; i++ {
		char := string(word[i])
		if char != " " {
			counter[char]++
		}
	}
	fmt.Println("Counter: ", counter)
}

// 2. Initialize the map with the following pairs: {1: "one", 2: "two", 3: "three"} and print whether numbers from 1 to 5 exist in this map or not.
func checkExistenceOfFirstFiveNumbers() {
	firstFiveNumbers := map[int]string{1: "one", 2: "two", 3: "three"}
	for i := 1; i <= 5; i++ {
		_, doesExist := firstFiveNumbers[i]
		fmt.Printf("Number: %d, Exists in Map: %t\n", i, doesExist)
	}
}

func main() {
	// frequencyCounter("Hello World")
	checkExistenceOfFirstFiveNumbers()
}
