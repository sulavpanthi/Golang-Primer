package main

import "fmt"

// 1. Create a new slice and print it
func createANewSlice() {
	arr := []int{1, 2, 3, 4, 5}
	slice1 := arr[1:4]
	slice2 := arr[:3]
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)
}

// 2. Create a new slice with user input length and capacity
func createSliceFromUserInput() {
	var lenSlice, capSlice int
	fmt.Print("Enter the length of slice: ")
	fmt.Scanln(&lenSlice)
	fmt.Print("Enter the capacity of slice: ")
	fmt.Scanln(&capSlice)

	slice := make([]int, lenSlice, capSlice)
	fmt.Println("Length of slice created is:", len(slice))
	fmt.Println("Capacity of slice created is:", cap(slice))
}

// 3. Create a new slice from existing slice using copy then modify the latest copy and see how it affects both source and destination slices
func copySlice() {
	source := []int{1, 2, 3}
	destination := make([]int, len(source)+5)
	copy(destination, source)
	destination[0] = 11
	fmt.Println("Source", source)
	fmt.Println("Destination", destination, destination[len(destination)-1])
}

// 4. Use append to make changes in a slice beyond its capacity and see how it affects the underlying array
func updateSliceUsingAppend() {
	arr := []int{10, 20, 30}
	slice1 := arr[:2]
	slice2 := append(slice1, 40, 50)
	slice2[0] = 1
	fmt.Println("slice1", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2", slice2, len(slice2), cap(slice2))
	fmt.Println("initial slice", arr, len(arr), cap(arr))
}

// 5. Use index in slice make changes in a slice and see how it affects the underlying array
func updateSliceUsingIndex() {
	arr := []int{10, 20, 30, 40, 50}
	slice1 := arr[:4]
	slice2 := arr[:2]
	slice2[0] = 99
	fmt.Println("slice1", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2", slice2, len(slice2), cap(slice2))
	fmt.Println("initial slice", arr, len(arr), cap(arr))
}

func main() {
	// createANewSlice()
	// createSliceFromUserInput()
	copySlice()
	// updateSliceUsingAppend()
	// updateSliceUsingIndex()
}
