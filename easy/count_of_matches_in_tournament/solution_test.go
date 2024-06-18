package count_of_matches_in_tournament_test

import (
	"testing"

	solution "github.com/mfonism/leetcode_with_rainbow_devs/easy/count_of_matches_in_tournament"
)

func TestCountFailsForZeroTeams(t *testing.T) {
	t.Parallel()

	_, err := solution.Count(0)
	if err == nil {
		t.Error("want non-nil error, got nil")
	}	
}

func TestCountFailsForOneTeam(t *testing.T) {
	t.Parallel()

	_, err := solution.Count(1)
	if err == nil {
		t.Error("want non-nil error, got nil")
	}	
}

