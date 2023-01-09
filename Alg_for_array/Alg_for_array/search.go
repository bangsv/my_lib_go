package Alg_for_array

func Line_search(array []int, n int) int { // linear search
	for i := 0; i < len(array); i++ {
		if array[i] == n {
			return i // if found
		}
	}
	return -1 // if not found
}

// Binary_search is a function that uses binary search to search for a number in an array.
// It takes two arguments: array, which is the array to be searched; and n, which is the number to be searched for.
// It returns the index of the number in the array if the number is found, and -1 if the number is not found.

func Binary_search(array []int, n int) int { // binary search
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := (low + high) / 2
		if array[mid] == n {
			return mid
		} else if array[mid] > n {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func Block_search(array []int, n int) int { // block search
	block := 2
	for i := 0; i < len(array); i += block {
		if array[i] == n {
			return i
		} else if array[i] > n {
			for j := i - block; j < i; j++ {
				if array[j] == n {
					return j
				}
			}
		}
	}
	return -1
}

func Interpolation_search(array []int, n int) int { // interpolation search
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := low + (n-array[low])*(high-low)/(array[high]-array[low])
		if array[mid] == n {
			return mid
		} else if array[mid] > n {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
