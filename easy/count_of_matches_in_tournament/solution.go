package count_of_matches_in_tournament

import "fmt"

// Problem at https://leetcode.com/problems/count-of-matches-in-tournament/

func Count(teams int) (int, error) {
	var matches int

	if teams < 2 {
		return matches, fmt.Errorf("cannot play tournament with %d matches", teams)
	}

	return matches, nil
}
