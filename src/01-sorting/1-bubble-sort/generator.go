package main

import (
	"fmt"
	"math/rand"
)

// MakeRandomSlice creates a slice with `numItems` entries, uses a loop to fill in the values, and returns the slice.
func makeRandomSlice(numItems, max int) []int {
	a := []int{}

	for i := 0; i < numItems; i++ {
		a = append(a, rand.Intn(max))
	}

	return a
}

// PrintSlice takes two parameters: a slice of integers and an integer numItems. The function should print either the full slice or the first numItems entries, whichever is smaller. (We’ll use this to see just part of really big slices.) at most numItems items.
func printSlice(slice []int, numItems int) {

	numItemsToPrint := min(len(slice), numItems)

	for i := 0; i < numItemsToPrint; i++ {
		fmt.Println(slice[i])
	}

}

// checkSorted function that takes a slice as a parameter and loops through its entries to verify that it is sorted. It should print either “The slice is sorted” or “The slice is NOT sorted!” accordingly.
func checkSorted(slice []int) {

	for i := 0; i < len(slice); i++ {
		if i+1 == len(slice) {
			break
		} else {
			if slice[i+1] < slice[i] {
				fmt.Println("The slice is NOT sorted!")
				return
			}
		}
	}

	fmt.Println("The slice is sorted")

}

// bubbleSort sorts the slice.
func bubbleSort(slice []int) []int {

	revolutions := len(slice)

	for revs := 0; revs < revolutions; revs++ {

		for index, _ := range slice {
			if index+1 == len(slice) {
				break
			}
			valOne := slice[index]
			valTwo := slice[index+1]
			if valOne > valTwo {
				slice[index] = valTwo
				slice[index+1] = valOne
			}
		}
	}
	return slice
}
