package first_solution

func Subsequence(array []int, subsequence []int) bool {

	i, j := 0, 0

	for i < len(array) && j < len(subsequence) {
		if array[i] == subsequence[j] {
			j++
		}
		i++

	}

	return j == len(subsequence)
}
