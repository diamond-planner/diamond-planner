package bsm

import "testing"

func TestDetermineTableRow_NilOrEmpty(t *testing.T) {
	team := Team{Name: "Berlin Skylarks", ShortName: "BSK"}

	if got := DetermineTableRow(team, nil); got != nil {
		t.Fatalf("expected nil for nil table, got %+v", got)
	}

	tbl := &Table{Rows: []Row{}}
	if got := DetermineTableRow(team, tbl); got != nil {
		t.Fatalf("expected nil for empty table, got %+v", got)
	}
}

func TestDetermineTableRow_ByTeamName(t *testing.T) {
	team := Team{Name: "Berlin Skylarks", ShortName: "BSK"}
	row := Row{TeamName: "Berlin Skylarks", ShortTeamName: "SKY"}
	tbl := &Table{Rows: []Row{row}}

	got := DetermineTableRow(team, tbl)
	if got == nil || got.TeamName != "Berlin Skylarks" {
		t.Fatalf("expected to find row by TeamName, got %+v", got)
	}
}

func TestDetermineTableRow_ByShortTeamName(t *testing.T) {
	team := Team{Name: "Other Name", ShortName: "BSK"}
	row := Row{TeamName: "Berlin Skylarks", ShortTeamName: "BSK"}
	tbl := &Table{Rows: []Row{row}}

	got := DetermineTableRow(team, tbl)
	if got == nil || got.ShortTeamName != "BSK" {
		t.Fatalf("expected to find row by ShortTeamName, got %+v", got)
	}
}

func TestDetermineTableRow_NoMatch(t *testing.T) {
	team := Team{Name: "Berlin Skylarks", ShortName: "BSK"}
	row := Row{TeamName: "Other Team", ShortTeamName: "OTH"}
	tbl := &Table{Rows: []Row{row}}

	if got := DetermineTableRow(team, tbl); got != nil {
		t.Fatalf("expected nil when no matching row, got %+v", got)
	}
}
