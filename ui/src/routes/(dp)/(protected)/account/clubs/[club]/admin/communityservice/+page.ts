import type {PageLoad} from './$types';
import {client} from "$lib/dp/client.svelte.ts";
import type {ClubCommunityServiceRow} from "$lib/dp/types/ClubCommunityServiceRow.ts";

export const load: PageLoad = (async ({fetch, params, depends, parent}) => {
  const rows = client.send<ClubCommunityServiceRow>(`/api/clubs/${params.club}/admin/communityservice`, {
    fetch: fetch,
  });
  const parentData = await parent();

  depends("communityservice:admin");

  return {
    rows: rows,
    club: parentData.club
  };
});
