package main

import (
	"fmt"
	"math"
)

// 1. Declare a package-scoped variable isAdmin with the value true and print it, and also declare and print a title variable with the value "Learn Golang" in the main function scope.
var isAdmin = true

func printVariablesInScopes() {
	title := "Learn Golang"
	fmt.Println(isAdmin)
	fmt.Println(title)
}

// 2. Declare two variables, myName and myAge, with values "Tim" and 4, respectively, then using short declaration, declare a new variable title with the value "Golang" and reassign "Learn Go" and 40 to myName and myAge, and print all three variables in one line.
func updateAndPrintVariables() {
	myName, myAge := "Tim", 4
	title, myName, myAge := "Golang", "Learn Go", 40
	fmt.Println(title, myName, myAge)
}

// 3. Declare two variables, one and two, with values "Golang" and "Python" respectively, swap their values, and print the swapped variables.
func swapAndPrintValues() {
	one, two := "Golang", "Python"
	one, two = two, one
	fmt.Println("One: ", one)
	fmt.Println("Two: ", two)
}

// 4. Declare three variables: name with type string, age with type int, and isActive with type boolean. Initialize them with their default values (zero values).
func initializeAndPrintDefaultValues() {
	var (
		name     string
		age      int
		isActive bool
	)
	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("IsActive: ", isActive)
}

// 5. In one line of code, declare three constants, daysPerWeek, daysPerMonth, and daysPerYear, and assign them values of 7, 30, and 365 respectively.
func declareAndPrintConstants() {
	const daysPerWeek, daysPerMonth, daysPerYear = 7, 30, 365
	fmt.Println(daysPerWeek, daysPerMonth, daysPerYear)
}

// 6. Declare two variables, p64 and p32 with each respective types float64 and float32, both assigned the value math.Pi, and print them on separate lines.
func printPiInDifferentFormats() {
	var (
		p64 float64 = math.Pi
		p32 float32 = math.Pi
	)
	fmt.Println("P64 ", p64)
	fmt.Println("P32 ", p32)
}

// 7. Declare and print a variable called daysPerYear with type uint16 and assign it the value 365.
func declareAndPrintDaysPerYear() {
	var daysPerYear uint16 = 365
	fmt.Println(daysPerYear)
}

// 8. Declare and print a variable toDecimal with type uint16 that takes the hexadecimal value "7e8" and converts it to decimal.
func convertHexToDecimalAndPrint() {
	var toDecimal uint16 = 0x7e8
	fmt.Printf("%d", toDecimal)
	fmt.Println()
}

// 9. Write a program to declare a float32 variable `x` with the value 300.345, convert it to an int, assign it to `y`, and print `y`.
func convertFloatToIntAndPrint() {
	var x float32 = 300.345
	y := int(x)
	fmt.Println(y)
}

// 10. Create a program that calculates the average weight of three people and prints the result to the screen.
func calculateAndPrintAverageWeight() {
	var person1, person2, person3 float32
	fmt.Print("Enter weight of first person: ")
	fmt.Scanln(&person1)
	fmt.Print("Enter weight of second person: ")
	fmt.Scanln(&person2)
	fmt.Print("Enter weight of third person: ")
	fmt.Scanln(&person3)
	avgWeight := (person1 + person2 + person3) / 3.0
	fmt.Println("Average weight is ", avgWeight)
}

// 11. Replace the question mark in each operation with an appropriate logical operator to ensure the result is always true in all operations.
func demonstrateLogicalOperators() {
	a := 100
	b := 200

	fmt.Println(b > a)
	fmt.Println(!(a == b && b == a))
	fmt.Println((a < b) || a == b)
	fmt.Println(b > a || a > b)
	fmt.Println(!(a > b))
	fmt.Println(a > b || b > a)
	fmt.Println(b > a)
}

// 12. Write a program that takes an exam score as input from the user and displays a message based on the score, printing "Congratulations, you are a genius!" for scores above 90, "Good job" for scores between 70 and 90, "Do better next time" for scores between 50 and 70, and "Try next time" for scores below 50.
func provideScoreFeedback() {
	var score int
	fmt.Print("Enter your academic score: ")
	fmt.Scanln(&score)
	switch {
	case score > 90:
		fmt.Println("Congratulations, you are a genius!")
	case 70 < score && score <= 90:
		fmt.Println("Good job")
	case 50 < score && score <= 70:
		fmt.Println("Do better next time")
	case score <= 50:
		fmt.Println("Try next time")
	}
}

// 13. Print pattern for number from 1 to 4 with occurrence count equal to their value
func printOccurrencePattern() {
	for i := 0; i <= 4; i++ {
		for x := 0; x < i; x++ {
			fmt.Print(i, "\t")
		}
		fmt.Println()
	}
}

// 14. Use an array to store values: "Golang", "Python", "Rust", "JavaScript", and "TypeScript". Print this array to the screen.
func printProgrammingLanguages() {
	languages := [5]string{"Golang", "Python", "Rust", "JavaScript", "TypeScript"}
	fmt.Println(languages)
}

func main() {
	// printVariablesInScopes()
	// updateAndPrintVariables()
	// swapAndPrintValues()
	// initializeAndPrintDefaultValues()
	// declareAndPrintConstants()
	// printPiInDifferentFormats()
	// declareAndPrintDaysPerYear()
	// convertHexToDecimalAndPrint()
	// convertFloatToIntAndPrint()
	// calculateAndPrintAverageWeight()
	// demonstrateLogicalOperators()
	// provideScoreFeedback()
	// printOccurrencePattern()
	printProgrammingLanguages()
}
