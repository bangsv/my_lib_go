package Alg_for_array

import (
	"math/rand"
	"time"
)

// The Sum function takes an array of integers and returns the sum of all the numbers in the array.
// The array is defined in the main function and passed to the Sum function.
func Sum(array []int) int {
	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}

// This code sets the random seed for the random number generator
// used by the Random function.
// The seed is set based on the current time.
// The seed is set only once at the beginning of the program.

func Random(array []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(100)
	}
}

// Delete function removes an element at a given index from a given array
// and returns the resulting array.
func Delete(array []int, index int) []int {
	array = append(array[:index], array[index+1:]...)
	return array
}
