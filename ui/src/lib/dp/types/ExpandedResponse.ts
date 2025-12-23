import type {RecordModel} from "pocketbase";
import type {UserParticipationDTO} from "$lib/dp/types/UserParticipationDTO.ts";
import type {
  AnnouncementsResponse,
  ClubsResponse,
  CommentsResponse,
  EventsCreate,
  EventseriesCreate,
  EventsResponse,
  LocationsResponse,
  ParticipationsResponse,
  ServiceentriesResponse,
  TeamsResponse,
  UniformsetsResponse,
  UsersResponse,
} from "./pb-types.ts";

export type Extension<T, E> = T & E

export type ExpandedEvent = Extension<
    EventsResponse,
    {
      expand: {
        participations_via_event?: Extension<
            ParticipationsResponse,
            {
              expand: {
                user: UsersResponse
              }
            }
        >[]
        attire?: UniformsetsResponse
        team?: ExpandedTeam
        location?: LocationsResponse
        comments_via_event?: ExpandedComment[]
      }
      participations: {
        in: ExpandedParticipation[]
        out: ExpandedParticipation[]
        maybe: ExpandedParticipation[]
        unspecified: UserParticipationDTO[]
      }
      userParticipation?: ExpandedParticipation
    }
>

export type ExpandedTeam = Extension<
    TeamsResponse,
    {
      expand: {
        club: ClubsResponse
        admins?: UsersResponse[]
      }
    }
>

export type ExpandedClub = Extension<
    ClubsResponse,
    {
      bsm_api_key: string
      expand: {
        admins: UsersResponse[]
      }
    }
>

export type ExpandedUniformSet = Extension<
    UniformsetsResponse,
    {
      expand?: {
        club?: ClubsResponse
      }
    }
>

export type ExpandedParticipation = Extension<
    ParticipationsResponse,
    {
      expand?: {
        user?: CustomAuthModel
      }
    }
>

export type CustomAuthModel = Extension<
    RecordModel,
    {
      email?: string
    } & UsersResponse
>

export type EventType = EventsResponse["type"]
export type ParticipationType = ParticipationsResponse["state"]

export type EventSeriesCreationData = Extension<
    EventseriesCreate,
    {
      location: EventsCreate["location"]
      desc: EventsCreate["desc"]
      series_start: string
      series_end: string
    }
>

export type ExpandedAnnouncement = Extension<
    AnnouncementsResponse,
    {
      expand?: {
        author: CustomAuthModel
        club?: ClubsResponse
        team?: TeamsResponse
        comments_via_announcement?: ExpandedComment[]
      }
    }
>

export type ExpandedComment = Extension<
    CommentsResponse,
    {
      expand?: {
        user?: CustomAuthModel
      }
    }
>

export type ExpandedServiceEntry = Extension<
  ServiceentriesResponse, {
  expand?: {
    club?: ClubsResponse
  }
}>
