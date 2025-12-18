import type {Gameday} from "bsm.js";

export type AppPreferences = {
  selectedSeason: number
  gameday: Gameday
  leagueGroupID: number
  favoriteTeamID: number
}
