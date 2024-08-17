package first_solution

import "testing"

func TestTournamentWinner(t *testing.T) {
	competitions := [][]string{
		{"TeamA", "TeamB"},
		{"TeamC", "TeamD"},
		{"TeamC", "TeamB"},
	}

	results := []int{0, 1, 0}

	expectedWinner := "TeamB"

	actualWinner := TournamentWinner(competitions, results)

	if actualWinner != expectedWinner {
		t.Errorf("Expected winner: %s, but got: %s", expectedWinner, actualWinner)
	}
}
