package main

func Sum(numbers [5]int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}

	return sum
}
