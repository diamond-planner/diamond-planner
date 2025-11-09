package bsm

import "testing"

func pint(i int) *int { return &i }

func newMatch(home, away string, homeRuns, awayRuns *int) Match {
	return Match{
		HomeTeamName: home,
		AwayTeamName: away,
		HomeRuns:     homeRuns,
		AwayRuns:     awayRuns,
	}
}

func TestPlayoffSeries_NoGames(t *testing.T) {
	ps := NewPlayoffSeries([]Match{})
	ps.GetSeriesStatus()

	if ps.Status != SeriesStatusTied {
		t.Errorf("expected status tied for empty series, got %v", ps.Status)
	}
	if ps.SeriesLength != 0 {
		t.Errorf("expected series length 0, got %d", ps.SeriesLength)
	}
	if len(ps.Teams) != 0 {
		t.Errorf("expected no teams initialized, got %d", len(ps.Teams))
	}
	if ps.LeadingTeam != nil || ps.TrailingTeam != nil {
		t.Errorf("expected leading and trailing to be nil, got %v and %v", ps.LeadingTeam, ps.TrailingTeam)
	}
}

func TestPlayoffSeries_SingleGame_HomeWins(t *testing.T) {
	hr, ar := 5, 3
	g := newMatch("TeamB", "TeamA", pint(hr), pint(ar))
	ps := NewPlayoffSeries([]Match{g})
	ps.GetSeriesStatus()

	if ps.SeriesLength != 1 {
		t.Fatalf("expected series length 1, got %d", ps.SeriesLength)
	}
	if ps.Status != SeriesStatusWon {
		t.Fatalf("expected status won for best-of-1, got %v", ps.Status)
	}
	if ps.LeadingTeam == nil || ps.LeadingTeam.Name != "TeamB" || ps.LeadingTeam.Wins != 1 {
		t.Fatalf("expected leading TeamB with 1 win, got %+v", ps.LeadingTeam)
	}
	if ps.TrailingTeam == nil || ps.TrailingTeam.Name != "TeamA" || ps.TrailingTeam.Wins != 0 {
		t.Fatalf("expected trailing TeamA with 0 wins, got %+v", ps.TrailingTeam)
	}
	if ps.Teams["TeamB"].Wins != 1 || ps.Teams["TeamA"].Wins != 0 {
		t.Fatalf("unexpected team wins, got B=%d A=%d", ps.Teams["TeamB"].Wins, ps.Teams["TeamA"].Wins)
	}
}

func TestPlayoffSeries_TwoGames_Split_Tied(t *testing.T) {
	g1 := newMatch("TeamB", "TeamA", pint(4), pint(2)) // TeamB wins
	g2 := newMatch("TeamB", "TeamA", pint(1), pint(2)) // TeamA (away) wins
	ps := NewPlayoffSeries([]Match{g1, g2})
	ps.GetSeriesStatus()

	if ps.SeriesLength != 2 {
		t.Fatalf("expected series length 2, got %d", ps.SeriesLength)
	}
	if ps.Status != SeriesStatusTied {
		t.Fatalf("expected status tied after split, got %v", ps.Status)
	}
	if ps.LeadingTeam != nil || ps.TrailingTeam != nil {
		t.Fatalf("expected no leading/trailing when tied, got %v/%v", ps.LeadingTeam, ps.TrailingTeam)
	}
	if ps.Teams["TeamA"].Wins != 1 || ps.Teams["TeamB"].Wins != 1 {
		t.Fatalf("unexpected team wins, got A=%d B=%d", ps.Teams["TeamA"].Wins, ps.Teams["TeamB"].Wins)
	}
}

func TestPlayoffSeries_TwoGames_Sweep_Won(t *testing.T) {
	g1 := newMatch("TeamB", "TeamA", pint(7), pint(1)) // TeamB wins
	g2 := newMatch("TeamB", "TeamA", pint(3), pint(0)) // TeamB wins
	ps := NewPlayoffSeries([]Match{g1, g2})
	ps.GetSeriesStatus()

	if ps.Status != SeriesStatusWon {
		t.Fatalf("expected status won for 2-0, got %v", ps.Status)
	}
	if ps.LeadingTeam == nil || ps.LeadingTeam.Name != "TeamB" || ps.LeadingTeam.Wins != 2 {
		t.Fatalf("expected leading TeamB with 2 wins, got %+v", ps.LeadingTeam)
	}
	if ps.TrailingTeam == nil || ps.TrailingTeam.Name != "TeamA" || ps.TrailingTeam.Wins != 0 {
		t.Fatalf("expected trailing TeamA with 0 wins, got %+v", ps.TrailingTeam)
	}
}

func TestPlayoffSeries_MixedUndecided_Leading(t *testing.T) {
	g1 := newMatch("TeamB", "TeamA", pint(2), pint(0)) // TeamB wins
	g2 := newMatch("TeamB", "TeamA", nil, nil)         // undecided
	g3 := newMatch("TeamB", "TeamA", nil, nil)         // undecided
	ps := NewPlayoffSeries([]Match{g1, g2, g3})
	ps.GetSeriesStatus()

	if ps.SeriesLength != 3 {
		t.Fatalf("expected series length 3, got %d", ps.SeriesLength)
	}
	if ps.Status != SeriesStatusLeading {
		t.Fatalf("expected status leading (not yet won), got %v", ps.Status)
	}
	if ps.LeadingTeam == nil || ps.LeadingTeam.Name != "TeamB" || ps.LeadingTeam.Wins != 1 {
		t.Fatalf("expected leading TeamB with 1 win, got %+v", ps.LeadingTeam)
	}
	if ps.TrailingTeam == nil || ps.TrailingTeam.Name != "TeamA" || ps.TrailingTeam.Wins != 0 {
		t.Fatalf("expected trailing TeamA with 0 wins, got %+v", ps.TrailingTeam)
	}
}

func TestPlayoffSeries_AllUndecided(t *testing.T) {
	g1 := newMatch("TeamB", "TeamA", nil, nil)
	g2 := newMatch("TeamB", "TeamA", nil, nil)
	ps := NewPlayoffSeries([]Match{g1, g2})
	ps.GetSeriesStatus()

	if ps.SeriesLength != 2 {
		t.Fatalf("expected series length 2, got %d", ps.SeriesLength)
	}
	if ps.Status != SeriesStatusTied {
		t.Fatalf("expected status tied when no wins, got %v", ps.Status)
	}
	if ps.LeadingTeam != nil || ps.TrailingTeam != nil {
		t.Fatalf("expected no leading/trailing when no wins, got %v/%v", ps.LeadingTeam, ps.TrailingTeam)
	}
	if len(ps.Teams) != 2 {
		t.Fatalf("expected teams initialized to 2, got %d", len(ps.Teams))
	}
	if ps.Teams["TeamA"].Wins != 0 || ps.Teams["TeamB"].Wins != 0 {
		t.Fatalf("expected both teams to have 0 wins, got A=%d B=%d", ps.Teams["TeamA"].Wins, ps.Teams["TeamB"].Wins)
	}
}
