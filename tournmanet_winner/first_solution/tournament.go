package first_solution

func TournamentWinner(competitions [][]string, results []int) string {

	scores := make(map[string]int)
	winner := ""

	for i := range competitions {
		localWinner := competitions[i][1-results[i]]
		scores[localWinner] += 3

		if scores[localWinner] > scores[winner] {
			winner = localWinner
		}

	}

	return winner
}
