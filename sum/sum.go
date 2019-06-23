package main

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return result
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		tails := numbers[1:]
		sums = append(sums, Sum(tails))
	}
	return
}
