package main

import (
	"fmt"
	"tournament_winner/first_solution"
)

func main() {
	competitions := [][]string{
		{"HTML", "C#"},
		{"C#", "Python"},
		{"Python", "HTML"},
	}
	results := []int{0, 0, 1}

	expectedWinner := "Python"
	result := first_solution.TournamentWinner(competitions, results)

	if result == expectedWinner {
		fmt.Println("Winner:", result)
	} else {
		fmt.Println("Wrong answer. Expected:", expectedWinner, "but got:", result)
	}
}
