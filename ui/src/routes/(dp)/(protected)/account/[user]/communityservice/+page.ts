import type {PageLoad} from './$types';
import {client} from "$lib/dp/client.svelte.ts";
import type {CommunityServiceData} from "$lib/dp/types/CommunityServiceData.ts";

export const load: PageLoad = (async ({params, url, fetch, depends}) => {
  let season = url.searchParams.get("season");
  if (!season) {
    season = new Date().getFullYear().toString();
  }

  const communityServiceData = client.send<CommunityServiceData>(`/api/communityservice/${params.user}/${season}`, {
    fetch: fetch,
    requestKey: `${params.user}-${season}-serviceentries`,
  });

  depends("communityservice:list");

  return {
    set: communityServiceData,
  };
});
