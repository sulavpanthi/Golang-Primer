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

// 5. Write a Go program that defines a function add which takes two integers and returns their sum.
func sum(a, b int) int {
	return a + b
}

// 6. Write a Go program that defines a function divide which takes two integers and returns both the quotient and the remainder.
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// 7. Write a Go program that defines a function rectangleAreaPerimeter which takes the length and width of a rectangle and returns both the area and the perimeter using named return values.
func rectangleAreaPerimeter(length, width int) (area, perimeter int) {
	area = length * width
	perimeter = 2 * (length + width)
	return
}

// 8. Create a variadic function sum that takes any number of elements, return their sum.
func variadicSum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	// frequencyCounter("Hello World")
	// checkExistenceOfFirstFiveNumbers()
	// sliceIterationUsingRange()
	// mapIterationUsingRange()
	// fmt.Println("Sum result: ", sum(5, 5))

	// quotient, remainder := divide(7, 5)
	// fmt.Printf("Quotient: %d, Remainder: %d\n", quotient, remainder)

	// area, perimeter := rectangleAreaPerimeter(4, 8)
	// fmt.Printf("Area: %d, Perimeter: %d\n", area, perimeter)

	fmt.Printf("Variadic sum of [1,2,3,4,5]: %d\n", variadicSum(1, 2, 3, 4, 5))
	fmt.Printf("Variadic sum of [11,21,31,41,51,61,71]: %d\n", variadicSum(11, 21, 31, 41, 51, 61, 71))

}
