package dp

import (
	"github.com/pocketbase/pocketbase/core"
)

const (
	LeagueGroupsCollection = "leaguegroups"
)

var _ core.RecordProxy = (*LeagueGroup)(nil)

// LeagueGroup RecordProxy for collection `leaguegroups`.
// Provides type-safe struct access to league group records.
type LeagueGroup struct {
	core.BaseRecordProxy
}

func (lg *LeagueGroup) ID() string {
	return lg.Id
}

func (lg *LeagueGroup) BSMID() int {
	return lg.GetInt("bsm_id")
}

func (lg *LeagueGroup) SetBSMID(bsmID int) {
	lg.Set("bsm_id", bsmID)
}

func (lg *LeagueGroup) Season() int {
	return lg.GetInt("season")
}

func (lg *LeagueGroup) SetSeason(season int) {
	lg.Set("season", season)
}

func (lg *LeagueGroup) Acronym() string {
	return lg.GetString("acronym")
}

func (lg *LeagueGroup) SetAcronym(acronym string) {
	lg.Set("acronym", acronym)
}

func (lg *LeagueGroup) Name() string {
	return lg.GetString("name")
}

func (lg *LeagueGroup) SetName(name string) {
	lg.Set("name", name)
}

func (lg *LeagueGroup) Clubs() []string {
	return lg.GetStringSlice("clubs")
}

func (lg *LeagueGroup) SetClubs(clubs []string) {
	lg.Set("clubs", clubs)
}
