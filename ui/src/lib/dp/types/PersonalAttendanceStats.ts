import type {EventType, ParticipationType} from "$lib/dp/types/ExpandedResponse.ts";

export interface PersonalAttendanceStatsItem {
  season: number;
  attendanceTotals: AttendanceTotal[];
  participationTotals: ParticipationTotal[];
  totalPossibleEvents: number;
  type: EventType;
  values: ParticipationStatsByPerson[];
  teamName: string;
}

export type ParticipationStatsByPerson = {
  id: string
  lastName: string
  firstName: string
  type: Exclude<EventType, "">
  inCount: number
  outCount: number
  maybeCount: number
  totalCount: number
}

export interface AttendanceTotal {
  type: EventType;
  attended: number;
  total: number;
}

export interface ParticipationTotal {
  type: Exclude<ParticipationType, "">;
  amount: number;
}
