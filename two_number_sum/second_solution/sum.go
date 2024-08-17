package second

import "sort"

func TwoNumberSum(array []int, target int) []int {
	sort.Ints(array)
	left, right := 0, len(array)-1

	for i := 0; i < len(array); i++ {

		sum := array[left] + array[right]

		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			return []int{array[left], array[right]}
		}
	}

	return []int{}
}
