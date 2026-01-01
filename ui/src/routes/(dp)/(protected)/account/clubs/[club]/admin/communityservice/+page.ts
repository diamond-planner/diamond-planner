import type {PageLoad} from './$types';
import {client} from "$lib/dp/client.svelte.ts";
import type {ClubCommunityServiceRow} from "$lib/dp/types/ClubCommunityServiceRow.ts";

export const load: PageLoad = (async ({fetch, params, url, depends, parent}) => {
  type Query = {
    season?: string;
  }
  const query: Query = {};

  let season = url.searchParams.get("season");
  if (!season) {
    season = new Date().getFullYear().toString();
  }
  query.season = season;

  const rows = client.send<ClubCommunityServiceRow[]>(`/api/clubs/${params.club}/admin/communityservice`, {
    fetch: fetch,
    query: query,
  });
  const parentData = await parent();

  depends("communityservice:admin");

  return {
    rows: rows,
    club: parentData.club
  };
});
