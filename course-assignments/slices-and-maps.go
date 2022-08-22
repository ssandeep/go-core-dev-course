package main

import "fmt"

// Helper functions:
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

//Slices

//1. Add given int(N) to each of []int elements.
func add_n_to_elements(n int, arr []int) {
	fmt.Println("1. Add given int(N) to each of []int elements.")
	fmt.Println("Original arr:", arr, "n:", n)
	for i, v := range arr {
		arr[i] = v + n
	}
	fmt.Println("The new arr is:", arr)
}

//2. Add the number int(N) to the end of the slice.
func append_n_to_slice(n int, arr []int) {
	fmt.Println("2. Add the number int(N) to the end of the slice.")

	fmt.Println("Before:")
	printSlice(arr)

	arr = append(arr, n)

	fmt.Println("After:")
	printSlice(arr)
}

//3. Add the number int(N) to the beginning of the slice.
func prepend_n_to_slice(n int, arr []int) {
	fmt.Println("3. Add the number int(N) to the beginning of the slice.")

	fmt.Println("Before:")
	printSlice(arr)

	arr = append(arr, n)

	fmt.Println("After:")
	printSlice(arr)
}

func main() {
	fmt.Printf("\nExercises for Slices to follow...\n\n")

	//1. Add given int(N) to each of []int elements.
	arr := []int{1, 2, 3, 4, 5}
	n := 10
	add_n_to_elements(n, arr)

	//2. Add the number int(N) to the end of the slice.
	append_n_to_slice(n, arr)

	//3. Add the number int(N) to the beginning of the slice.
	prepend_n_to_slice(n, arr)
}
