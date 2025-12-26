import type {ClubsResponse, ServiceentriesResponse, UsersResponse} from "$lib/dp/types/pb-types.ts";

export type CommunityServiceItem = {
  club: ClubsResponse
  current_minutes: number
  entries: ServiceentriesResponse[]
}

export type CommunityServiceData = {
  user: UsersResponse,
  season: number
  items: CommunityServiceItem[]
}
