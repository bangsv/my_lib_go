package Alg_for_array

// All sorted
func Bubble_sort(array []int) { // bubble sort
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

func Insertion_sort(array []int) { // insertion sort
	for i := 0; i < len(array); i++ {
		for j := i; j > 0; j-- {
			if array[j] < array[j-1] {
				array[j], array[j-1] = array[j-1], array[j]
			}
		}
	}
}

func Selection_sort(array []int) { // selection sort
	for i := 0; i < len(array); i++ {
		min := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		array[i], array[min] = array[min], array[i]
	}
}

func Merge(left, right []int) []int { // merge sort
	var result []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	result = append(result, left...)
	result = append(result, right...)
	return result
}

func Merge_sort(array []int) []int { // merge sort
	if len(array) < 2 {
		return array
	}
	mid := len(array) / 2
	left := Merge_sort(array[:mid])
	right := Merge_sort(array[mid:])
	return Merge(left, right)
}

func Quick_sort(array []int) []int { // quick sort
	if len(array) < 2 {
		return array
	}
	pivot := array[0]
	var left, right []int
	for i := 1; i < len(array); i++ {
		if array[i] < pivot {
			left = append(left, array[i])
		} else {
			right = append(right, array[i])
		}
	}
	return append(append(Quick_sort(left), pivot), Quick_sort(right)...)
}
