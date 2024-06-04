package main

import (
	"fmt"
	"subsequence/first_solution"
)

func main() {
	answer := first_solution.Subsequence([]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{1, 6, -1, 10})
	if answer {
		fmt.Println("Subsequence")
	} else {
		fmt.Println("Not a Subsequence")
	}

	answer = first_solution.Subsequence([]int{3, 1, 6, 7}, []int{3, 6, 7})
	if answer {
		fmt.Println("Subsequence")
	} else {
		fmt.Println("Not a Subsequence")
	}

	answer = first_solution.Subsequence([]int{2, 1, 4, 9, 0, 7}, []int{1, 7, 2})
	if answer {
		fmt.Println("Subsequence")
	} else {
		fmt.Println("Not a Subsequence")
	}
}
