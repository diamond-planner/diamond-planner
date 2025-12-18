import {client} from "$lib/dp/client.svelte.js";
import type {ExpandedClub} from "$lib/dp/types/ExpandedResponse.ts";
import type {LocationsResponse} from "$lib/dp/types/pb-types.js";
import type {PageLoad} from "../../../../../../../../.svelte-kit/types/src/routes";

export const load = (async ({fetch, params, depends}) => {
  const club = await client.collection("clubs").getOne<ExpandedClub>(params.club, {
    expand: "admins",
    fetch: fetch,
  });

  const locations = await client.collection("locations").getFullList<LocationsResponse>({
    filter: `club.id = "${params.club}"`,
    fetch: fetch,
  });

  depends("club:locations");

  return {
    club: club,
    locations: locations,
  };
}) satisfies PageLoad;
