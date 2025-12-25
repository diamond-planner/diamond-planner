import type {ClubsResponse, ServiceentriesResponse, UsersResponse} from "$lib/dp/types/pb-types.ts";

export type CommunityServiceItem = {
  club: ClubsResponse
  entries: ServiceentriesResponse
}

export type CommunityServiceData = {
  user: UsersResponse,
  season: number
  items: CommunityServiceItem[]
}
