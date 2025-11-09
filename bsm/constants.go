package bsm

const (
	LeagueGroupsCollection = "leaguegroups"
)

// BSM API filters

const (
	LeagueFilter  = "filters[leagues][]"
	SearchFilter  = "search"
	SeasonFilter  = "filters[seasons][]"
	GamedayFilter = "filters[gamedays][]"
	CompactFilter = "compact"
)

const (
	TypeBatting  = "batting"
	TypePitching = "pitching"
	TypeFielding = "fielding"
)

const TimeFormat = "2006-01-02 15:04:05 -0700"

// Batting constants
const (
	Runs               = "runs"
	RunsBattedIn       = "runs_batted_in"
	Hits               = "hits"
	Doubles            = "doubles"
	Triples            = "triples"
	Homeruns           = "homeruns"
	Strikeouts         = "strikeouts"
	BaseOnBalls        = "base_on_balls"
	HitByPitches       = "hit_by_pitches"
	StolenBases        = "stolen_bases"
	CaughtStealings    = "caught_stealings"
	BattingAverage     = "batting_average"
	OnBasePercentage   = "on_base_percentage"
	SluggingPercentage = "slugging_percentage"
	OnBasePlusSlugging = "on_base_plus_slugging"
)

// Fielding constants
const (
	Assists         = "assists"
	Putouts         = "putouts"
	Errors          = "errors"
	FieldingAverage = "fielding_average"
	DoublePlays     = "double_plays"
	TriplePlays     = "triple_plays"
	PassedBalls     = "passed_balls"
)

// Pitching constants
const (
	Games                         = "games"
	CompleteGames                 = "complete_games"
	InningsPitched                = "innings_pitched"
	BattersFaced                  = "batters_faced"
	PitchingRuns                  = "runs"
	EarnedRuns                    = "earned_runs"
	PitchingHits                  = "hits"
	PitchingDoubles               = "doubles"
	PitchingTriples               = "triples"
	PitchingHomeruns              = "homeruns"
	PitchingStrikeouts            = "strikeouts"
	BaseOnBallsAllowed            = "base_on_balls_allowed"
	HitByPitchesAllowed           = "hit_by_pitches"
	WildPitches                   = "wild_pitches"
	Wins                          = "wins"
	Losses                        = "losses"
	Saves                         = "saves"
	EarnedRunAverage              = "earned_runs_average"
	WalksAndHitsPerInningsPitched = "walks_and_hits_per_innings_pitched"
)

func GetBattingStatsTypes() []string {
	return []string{
		Runs,
		RunsBattedIn,
		Hits,
		Doubles,
		Triples,
		Homeruns,
		Strikeouts,
		BaseOnBalls,
		HitByPitches,
		StolenBases,
		CaughtStealings,
		BattingAverage,
		OnBasePercentage,
		SluggingPercentage,
		OnBasePlusSlugging,
	}
}

func GetFieldingStatsTypes() []string {
	return []string{
		Assists,
		Putouts,
		Errors,
		FieldingAverage,
		DoublePlays,
		TriplePlays,
		PassedBalls,
	}
}

func GetPitchingStatsTypes() []string {
	return []string{
		Games,
		CompleteGames,
		InningsPitched,
		BattersFaced,
		PitchingRuns,
		EarnedRuns,
		PitchingHits,
		PitchingDoubles,
		PitchingTriples,
		PitchingHomeruns,
		PitchingStrikeouts,
		BaseOnBallsAllowed,
		HitByPitchesAllowed,
		WildPitches,
		Wins,
		Losses,
		Saves,
		EarnedRunAverage,
		WalksAndHitsPerInningsPitched,
	}
}
