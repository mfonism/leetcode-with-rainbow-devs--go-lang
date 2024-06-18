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

func TestCountForTwoTeamsReturnsOneMatch(t *testing.T) {
	t.Parallel()

	c, err := solution.Count(2)

	if err != nil {
		t.Error("want nil error, got non-nil")
	}

	if c != 1 {
		t.Errorf("want 1, got %d", c)
	}

}

func TestCountForThreeTeamsReturnsTwoMatches(t *testing.T) {
	t.Parallel()

	teams := 3
	want := 2
	got, err := solution.Count(teams)

	if err != nil {
		t.Error("want nil error, got non-nil")
	}

	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}

}

func TestCountForFourTeamsReturnsThreeMatches(t *testing.T) {
	t.Parallel()

	teams := 4
	want := 3
	got, err := solution.Count(teams)

	if err != nil {
		t.Error("want nil error, got non-nil")
	}

	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}

}

func TestCountForSevenTeamsReturnsSixMatches(t *testing.T) {
	t.Parallel()

	teams := 7
	want := 6
	got, err := solution.Count(teams)

	if err != nil {
		t.Error("want nil error, got non-nil")
	}

	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}

}

func TestCountForFourteenTeamsReturnsThirteenMatches(t *testing.T) {
	t.Parallel()

	teams := 14
	want := 13
	got, err := solution.Count(teams)

	if err != nil {
		t.Error("want nil error, got non-nil")
	}

	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}

}
