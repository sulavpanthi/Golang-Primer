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

// 3. Creates a slice of integers: []int{10, 20, 30, 40, 50}. Use the range keyword to iterate over the slice and print each index and value.
func sliceIterationUsingRange() {
	nums := []int{10, 20, 30, 40, 50}
	for index, num := range nums {
		fmt.Printf("Index: %d, Number: %d\n", index, num)
	}
}

// 4. Creates a map with string keys and float values, e.g., map[string]float64{"pi": 3.14, "e": 2.71}. Use range to iterate over the map and print each key and value.
func mapIterationUsingRange() {
	constants := map[string]float64{"pi": 3.14, "e": 2.71}
	for key, val := range constants {
		fmt.Printf("Key: %s, Value: %f\n", key, val)
	}
}

func main() {
	// frequencyCounter("Hello World")
	// checkExistenceOfFirstFiveNumbers()
	// sliceIterationUsingRange()
	mapIterationUsingRange()
}
