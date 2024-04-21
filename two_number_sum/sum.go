package main

func TwoNumberSum(array []int, target int) []int {

	var firstNumber, secondNumber, sum int

	for i := 0; i < len(array)-1; i++ {
		firstNumber = array[i]
		for j := i + 1; j < len(array); j++ {
			secondNumber = array[j]
			sum = firstNumber + secondNumber
			if sum == target {
				return []int{firstNumber, secondNumber}
			}
		}
	}

	return []int{}
}
