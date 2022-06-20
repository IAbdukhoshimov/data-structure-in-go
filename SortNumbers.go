package main

import "sort"

func SortNumbers(numbers []int) []int {
	if len(numbers) > 0 {

		sort.Ints(numbers)
	}
	return numbers
}
