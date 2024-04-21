package main

import (
	"fmt"
	"reflect"
)

func main() {
	exampleNumberList := []int{1, 3, 6, -12, 3, 0, -5, 2, 4, 7}
	target := -10
	answer := TwoNumberSum(exampleNumberList, target)
	if reflect.DeepEqual(answer, []int{-12, 2}) {
		fmt.Println("Correct answer")
	} else {
		fmt.Println("Wrong answer")
	}
}
