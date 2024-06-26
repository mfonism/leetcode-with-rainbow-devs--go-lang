package count_of_matches_in_tournament

import "fmt"

// Problem at https://leetcode.com/problems/count-of-matches-in-tournament/

// Simple interface for LeetCode
func numberOfMatches(teams int) int {
	res, _ := Count(teams)

	return res
}

func Count(teams int) (int, error) {
	var matches int

	if teams < 1 {
		return matches, fmt.Errorf("cannot play tournament with %d teams", teams)
	}

	return doCount(teams), nil
}

func doCount(teams int) int {
	if teams <= 2 {
		return 1
	}

	matchesThisRound, restingTeam := teams/2, teams%2
	return matchesThisRound + doCount(matchesThisRound+restingTeam)
}
