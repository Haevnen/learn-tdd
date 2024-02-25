package main

func Sum(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}

	return sum
}

func SumAll(inputs ...[]int) []int {
	var res []int
	for _, numbers := range inputs {
		res = append(res, Sum(numbers))
	}

	return res
}

func SumAllTails(inputs ...[]int) []int {
	var res []int
	for _, numbers := range inputs {
		if len(numbers) == 0 {
			res = append(res, 0)
		} else {
			res = append(res, Sum(numbers[1:]))
		}
	}
	return res
}
