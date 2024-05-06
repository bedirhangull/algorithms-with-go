package second

import (
	"testing"
)

func TestTwoNumberSum(t *testing.T) {

	exampleNumberList := []int{2, 3, -1, -4, 10, 15}
	target := 11
	answer := TwoNumberSum(exampleNumberList, target)

	if answer[0] != -4 || answer[1] != 15 {
		t.Errorf("expectation -4,15, but got %v", answer)
	}

	exampleNumberList = []int{}
	target = 3
	answer = TwoNumberSum(exampleNumberList, target)

	if len(answer) != 0 {
		t.Errorf("expectation [], %v", answer)
	}

	exampleNumberList = []int{2, 1, 5, -4, 9, -10, 4}
	target = 6
	answer = TwoNumberSum(exampleNumberList, target)

	if answer[0] != 1 || answer[1] != 5 {
		t.Errorf("expectation 1,5, but got %v", answer)
	}

}
